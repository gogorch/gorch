package engine

import (
	"time"

	"github.com/gogorch/gorch/internal/lang/ast"
)

type args struct {
	argMap map[string]ArgValue
}

func (a *args) Arg(key string) ArgValue {
	v, has := a.argMap[key]
	if has {
		return v
	}

	return defaultArg
}

func (a *args) Has(key string) bool {
	_, has := a.argMap[key]
	return has
}

func fromArgsStmt(a *ast.ArgsStmt) *args {
	if a == nil {
		return &args{make(map[string]ArgValue)}
	}
	argMap := make(map[string]ArgValue)

	for _, arg := range a.ArgList {
		argMap[arg.Key] = fromArgStmt(arg)
	}
	return &args{argMap}
}

// ArgValue 参数值接口
type ArgValue interface {
	Int64(...int) int64
	Bool(...int) bool
	Str(...int) string
	Duration(...int) time.Duration

	Int64List() []int64
	BoolList() []bool
	StrList() []string
	DurationList() []time.Duration
}

type arg struct {
	Type ast.ConstantType

	IntegerValue  []int64
	StringValue   []string
	BooleanValue  []bool
	DurationValue []time.Duration
}

func fromArgStmt(a *ast.ArgStmt) *arg {
	arg := &arg{}
	arg.Type = a.Type
	if a.Type == ast.StringType {
		arg.Type = ast.StringType
		arg.StringValue = make([]string, len(a.StringValue))
		copy(arg.StringValue, a.StringValue)
	} else if a.Type == ast.BooleanType {
		arg.Type = ast.BooleanType
		arg.BooleanValue = make([]bool, len(a.BooleanValue))
		copy(arg.BooleanValue, a.BooleanValue)
	} else if a.Type == ast.IntegerType {
		arg.Type = ast.IntegerType
		arg.IntegerValue = make([]int64, len(a.IntegerValue))
		copy(arg.IntegerValue, a.IntegerValue)
	} else if a.Type == ast.DurationType {
		arg.Type = ast.DurationType
		arg.DurationValue = make([]time.Duration, len(a.DurationValue))
		copy(arg.DurationValue, a.DurationValue)
	}

	return arg
}

func (c *arg) Bool(tidx ...int) bool {
	if c.Type == ast.BooleanType {
		idx := 0
		if len(tidx) == 1 {
			idx = tidx[0]
		}
		if len(c.BooleanValue) > idx {
			return c.BooleanValue[idx]
		}
	}
	return false
}

func (c *arg) Int64(tidx ...int) int64 {
	if c.Type == ast.IntegerType {
		idx := 0
		if len(tidx) == 1 {
			idx = tidx[0]
		}
		if len(c.IntegerValue) > idx {
			return c.IntegerValue[idx]
		}
	}
	return 0
}

func (c *arg) Str(tidx ...int) string {
	if c.Type == ast.StringType {
		idx := 0
		if len(tidx) == 1 {
			idx = tidx[0]
		}
		if len(c.StringValue) > idx {
			return c.StringValue[idx]
		}
	}
	return ""
}

func (c *arg) Duration(tidx ...int) time.Duration {
	if c.Type == ast.DurationType {
		idx := 0
		if len(tidx) == 1 {
			idx = tidx[0]
		}
		if len(c.DurationValue) > idx {
			return c.DurationValue[idx]
		}
	}
	return 0
}

func (c *arg) BoolList() []bool {
	return c.BooleanValue
}

func (c *arg) Int64List() []int64 {
	return c.IntegerValue
}

func (c *arg) StrList() []string {
	return c.StringValue
}

func (c *arg) DurationList() []time.Duration {
	return c.DurationValue
}

var (
	defaultArg = &emptyArg{}
)

type emptyArg struct {
}

func (a *emptyArg) Bool(...int) bool {
	return false
}

func (a *emptyArg) Int64(...int) int64 {
	return 0
}

func (a *emptyArg) Str(...int) string {
	return ""
}

func (a *emptyArg) Duration(...int) time.Duration {
	return 0
}

func (a *emptyArg) BoolList() []bool {
	return []bool{}
}

func (a *emptyArg) Int64List() []int64 {
	return []int64{}
}

func (a *emptyArg) StrList() []string {
	return []string{}
}

func (a *emptyArg) DurationList() []time.Duration {
	return []time.Duration{}
}
