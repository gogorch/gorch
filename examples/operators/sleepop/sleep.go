package sleepop

import (
	"errors"
	"time"

	"github.com/gogorch/gorch"
)

// SleepOp 睡眠算子，在测试时睡眠指定时间
type SleepOp struct{}

func (s *SleepOp) Execute(ctx *gorch.Context) error {
	to := ctx.Arg("sleep").Duration()
	if to <= 0 {
		return errors.New("sleep operator must get sleep argument")
	}

	time.Sleep(to)
	return nil
}
