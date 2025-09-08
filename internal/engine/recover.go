package engine

import (
	"bytes"
	"fmt"
	"runtime"

	"github.com/gogorch/gorch/mlog"
)

// RecoverPanic 函数用于从 panic 中恢复，并返回包含错误信息和堆栈跟踪的日志属性切片。
// 参数：
//   - msg：触发 panic 时的错误信息，可以是任何类型。
//
// 返回值：
//   - []slog.Attr：包含错误信息和堆栈跟踪的日志属性切片。
func RecoverPanic(msg any) []mlog.LogValue {
	trace := make([]byte, 4096)
	n := runtime.Stack(trace[:], false)

	// 过滤掉 cover.go 的堆栈行
	lines := bytes.Split(trace[:n], []byte("\n"))
	filtered := make([][]byte, 0, len(lines))
	for _, line := range lines {
		if !bytes.Contains(line, []byte("recover.go")) {
			filtered = append(filtered, line)
		}
	}
	cleanStack := bytes.Join(filtered, []byte("\n"))

	return []mlog.LogValue{
		mlog.String("panic", fmt.Sprint(msg)),
		mlog.String("stack", string(cleanStack)),
	}
}
