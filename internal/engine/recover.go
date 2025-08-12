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

	return []mlog.LogValue{
		mlog.String("executePanic", fmt.Sprint(msg)),
		mlog.String("stack", string(bytes.ReplaceAll(trace[:n], []byte("\n"), []byte("\\n")))),
	}
}
