package recorder

import (
	"strconv"
	"sync"
	"time"

	"github.com/gogorch/gorch/pool"
)

// Recorder 算子执行结果记录器
type Recorder interface {
	// EnableLogOperatorName 记录算子名称到日志中
	EnableLogOperatorName()

	// CostThreshold 记录算子耗时的分界点，大于这个耗时就会记录算子耗时，否则算子不记录耗时
	// 你可以修改这个全局变量来改变行为
	SetCostThreshold(time.Duration)

	// StartRecording 开始记录引擎执行
	StartRecording()

	// StopRecording 引擎执行结束
	StopRecording()

	// RecordOperator 开始记录算子执行
	// 返回一个函数，当算子执行结束时，调用这个函数，记录算子执行结果
	RecordOperator(name string, seq string) func(int)

	// ElapsedTime 返回从开始执行到现在的时间
	ElapsedTime() time.Duration

	// GetOperatorRecords 返回算子执行记录的字符串
	// 字符串格式：算子名称,算子开始执行时间,算子执行耗时,算子执行状态
	// 如果算子执行成功，那么默认状态码为0，为0时不记录日志；否则状态码为算子执行返回的状态码
	// 如果算子的执行时间小于CostThreshold，那么不记录算子的执行时间
	// 默认CostThreshold为100微秒
	// 多个算子执行记录之间默认用 | 分隔
	GetOperatorRecords() string

	// GetTotalCost 返回引擎执行总耗时
	GetTotalCost() string

	// Release 回收对象到对象池
	Release()
}

type recorder struct {
	mutex sync.Mutex

	// 日志里记录执行的算子名称
	logOperatorName bool
	// 打印算子耗时的阈值，当算子的耗时大于这个阈值，才会记录算子的耗时到日志
	costThreshold time.Duration

	// 开始执行时间
	startTime time.Time
	// 总执行时间
	cost time.Duration

	// ors 算子执行记录
	ors []*operatorRecorder
}

type operatorRecorder struct {
	// name 算子名称
	name string
	// startOff 算子开始执行时间
	startOff time.Duration
	// cost 算子执行耗时
	cost time.Duration
	// status 算子执行返回的状态码
	status int
}

var (
	zeroTime = time.Time{}

	operatorRecorderSlicePool = pool.NewSlicePool[*operatorRecorder](0, 10240)

	recorderPool = pool.NewObjectPool(func(r *recorder) {
		r.startTime = zeroTime
		r.costThreshold = time.Microsecond * 100
		r.ors = operatorRecorderSlicePool.Get()
		r.ors = r.ors[:0]
	})

	operatorRecorderPool = pool.NewObjectPool(func(r *operatorRecorder) {
		r.name = ""
		r.startOff = 0
		r.cost = 0
		r.status = 0
	})
)

func NewRecorder() Recorder {
	r := recorderPool.Get()
	return r
}

func (r *recorder) StartRecording() {
	r.startTime = time.Now()
}

func (r *recorder) RecordOperator(name string, seq string) func(int) {
	start := time.Now()
	return func(i int) {
		cost := time.Since(start)
		or := operatorRecorderPool.Get()
		or.name = name
		if !r.logOperatorName {
			or.name = seq
		}
		or.startOff = time.Since(r.startTime)
		or.cost = cost
		or.status = i
		r.mutex.Lock()
		r.ors = append(r.ors, or)
		r.mutex.Unlock()
	}
}

func (r *recorder) EnableLogOperatorName() {
	r.logOperatorName = true
}

func (r *recorder) SetCostThreshold(t time.Duration) {
	r.costThreshold = t
}

func (r *recorder) ElapsedTime() time.Duration {
	return time.Since(r.startTime)
}

func (r *recorder) StopRecording() {
	r.cost = time.Since(r.startTime)
}

func (r *recorder) GetOperatorRecords() string {
	buf := pool.BytesBufferPool.Get()
	defer pool.BytesBufferPool.Put(buf)
	for i, or := range r.ors {
		buf.WriteString(or.name)
		if or.cost > r.costThreshold {
			buf.WriteString(",")
			buf.WriteString(durationToMillStr(or.startOff))
			buf.WriteString(",")
			buf.WriteString(durationToMillStr(or.cost))
		}
		if or.status != 0 {
			buf.WriteString(",")
			buf.WriteString(strconv.Itoa(or.status))
		}
		if i != len(r.ors)-1 {
			buf.WriteString("|")
		}
	}

	return buf.String()
}

// durationToMillStr 将 time.Duration 转换为毫秒的小数形式，并返回字符串
func durationToMillStr(d time.Duration) string {
	ms := float64(d) / 1e6                     // 将Duration转换为毫秒
	return strconv.FormatFloat(ms, 'f', 3, 64) // 格式化为字符串，保留三位小数
}

func (r *recorder) GetTotalCost() string {
	return durationToMillStr(r.cost)
}

func (r *recorder) Release() {
	if r == nil {
		return
	}

	for _, or := range r.ors {
		operatorRecorderPool.Put(or)
	}
	recorderPool.Put(r)
}
