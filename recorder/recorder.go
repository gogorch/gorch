package recorder

import (
	"strconv"
	"sync"
	"time"

	"github.com/gogorch/gorch/pool"
)

// Recorder 提供算子级执行记录能力
type Recorder interface {
	// WithOperatorName 设置是否在日志中使用算子名称而非序号
	UseOperatorName(bool) Recorder

	// Start 开始一次引擎执行记录
	Start()
	// Stop 结束当前引擎执行记录
	Stop()
	// Elapsed 返回当前已执行耗时
	Elapsed() time.Duration

	// Record 生成一个算子记录句柄，调用返回的 Done 函数即完成记录
	Record(name, seq string) Done

	// Report 生成本次执行报告
	Report() Report
	// Reset 回收对象到池，重置状态以便复用
	Reset()
}

// Done 用于完成单个算子记录
type Done func(status int)

// Report 提供只读的执行报告
type Report struct {
	TotalCost time.Duration
	Operators []operatorSnapshot
}

// OperatorRecord 单条算子记录
type OperatorRecord struct {
	Name     string
	StartOff time.Duration
	Cost     time.Duration
	Status   int
}

/* 实现部分 */

type recorder struct {
	logName bool
	start   time.Time
	total   time.Duration

	// 无锁化热点路径：先写本地缓存，再一次性合并
	localOps []operatorSnapshot
	mu       sync.Mutex
}

type operatorSnapshot struct {
	name     string
	startOff time.Duration
	cost     time.Duration
	status   int
}

var (
	recorderPool = pool.NewObjectPool(func(r *recorder) {
		r.localOps = make([]operatorSnapshot, 0, 128)
	})
)

// New 从池中获取一个 Recorder
func New() Recorder { return recorderPool.Get() }

func (r *recorder) UseOperatorName(on bool) Recorder {
	r.logName = on
	return r
}

func (r *recorder) Start() {
	r.start = time.Now()
}

func (r *recorder) Stop() {
	r.total = time.Since(r.start)
}

func (r *recorder) Elapsed() time.Duration {
	return time.Since(r.start)
}

func (r *recorder) Record(name, seq string) Done {
	start := time.Now()
	key := seq
	if r.logName {
		key = name
	}
	return func(status int) {
		snap := operatorSnapshot{
			name:     key,
			startOff: time.Since(r.start),
			cost:     time.Since(start),
			status:   status,
		}
		r.mu.Lock()
		r.localOps = append(r.localOps, snap)
		r.mu.Unlock()
	}
}

func (r *recorder) Report() Report {
	return Report{
		TotalCost: r.total,
		Operators: r.localOps,
	}
}

func (r *recorder) Reset() {
	// 回收前重置状态
	r.localOps = r.localOps[:0]
	r.total = 0
	r.start = time.Time{}
	recorderPool.Put(r)
}

// Format 将 Report 格式化为字符串，避免每次分配 bytes.Buffer
func (rep Report) Format(threshold time.Duration) string {
	buf := pool.BytesBufferPool.Get()
	defer pool.BytesBufferPool.Put(buf)

	for i, op := range rep.Operators {
		buf.WriteString(op.name)
		if op.cost > threshold {
			buf.WriteString(",")
			buf.WriteString(durationToMillStr(op.startOff))
			buf.WriteString(",")
			buf.WriteString(durationToMillStr(op.cost))
		}
		if op.status != 0 {
			buf.WriteString(",")
			buf.WriteString(strconv.Itoa(op.status))
		}
		if i != len(rep.Operators)-1 {
			buf.WriteString("|")
		}
	}
	return buf.String()
}

func durationToMillStr(d time.Duration) string {
	return strconv.FormatFloat(float64(d)/1000000, 'f', 2, 64) + "ms"
}
