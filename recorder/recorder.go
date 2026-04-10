package recorder

import (
	"strconv"
	"sync"
	"time"

	"github.com/gogorch/gorch/pool"
)

// Recorder 提供算子级执行记录能力
type Recorder interface {
	// UseOperatorName 设置是否在日志中使用算子名称而非序号
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
	Operators []OperatorRecord
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

	localOps []OperatorRecord
	mu       sync.Mutex
}

var (
	recorderPool = pool.NewObjectPool(func(r *recorder) {
		r.reset()
	})
)

// New 从池中获取一个 Recorder
func New() Recorder { return recorderPool.Get() }

func (r *recorder) UseOperatorName(on bool) Recorder {
	r.mu.Lock()
	r.logName = on
	r.mu.Unlock()
	return r
}

func (r *recorder) Start() {
	r.mu.Lock()
	r.start = time.Now()
	r.total = 0
	if r.localOps == nil {
		r.localOps = make([]OperatorRecord, 0, 128)
	} else {
		r.localOps = r.localOps[:0]
	}
	r.mu.Unlock()
}

func (r *recorder) Stop() {
	r.mu.Lock()
	r.total = time.Since(r.start)
	r.mu.Unlock()
}

func (r *recorder) Elapsed() time.Duration {
	r.mu.Lock()
	start := r.start
	r.mu.Unlock()
	return time.Since(start)
}

func (r *recorder) Record(name, seq string) Done {
	r.mu.Lock()
	base := r.start
	start := time.Now()
	key := seq
	if r.logName {
		key = name
	}
	r.mu.Unlock()

	startOff := time.Duration(0)
	if !base.IsZero() {
		startOff = start.Sub(base)
	}

	return func(status int) {
		rec := OperatorRecord{
			Name:     key,
			StartOff: startOff,
			Cost:     time.Since(start),
			Status:   status,
		}
		r.mu.Lock()
		r.localOps = append(r.localOps, rec)
		r.mu.Unlock()
	}
}

func (r *recorder) Report() Report {
	r.mu.Lock()
	defer r.mu.Unlock()

	ops := make([]OperatorRecord, len(r.localOps))
	copy(ops, r.localOps)
	return Report{
		TotalCost: r.total,
		Operators: ops,
	}
}

func (r *recorder) Reset() {
	r.mu.Lock()
	r.reset()
	r.mu.Unlock()
	recorderPool.Put(r)
}

func (r *recorder) reset() {
	r.logName = false
	r.start = time.Time{}
	r.total = 0
	if r.localOps == nil {
		r.localOps = make([]OperatorRecord, 0, 128)
	} else {
		r.localOps = r.localOps[:0]
	}
}

// Format 将 Report 格式化为字符串，避免每次分配 bytes.Buffer
func (rep Report) Format(threshold time.Duration) string {
	buf := pool.BytesBufferPool.Get()
	defer pool.BytesBufferPool.Put(buf)

	for i, op := range rep.Operators {
		buf.WriteString(op.Name)
		if op.Cost > threshold {
			buf.WriteString(",")
			buf.WriteString(durationToMillStr(op.StartOff))
			buf.WriteString(",")
			buf.WriteString(durationToMillStr(op.Cost))
		}
		if op.Status != 0 {
			buf.WriteString(",")
			buf.WriteString(strconv.Itoa(op.Status))
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
