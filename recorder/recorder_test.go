package recorder

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestReportOperatorRecord(t *testing.T) {
	r := New()
	r.Start()
	done := r.Record("op_name", "11")
	done(7)
	r.Stop()

	rep := r.Report()
	if assert.Len(t, rep.Operators, 1) {
		assert.Equal(t, "11", rep.Operators[0].Name)
		assert.Equal(t, 7, rep.Operators[0].Status)
		assert.GreaterOrEqual(t, rep.Operators[0].Cost, time.Duration(0))
	}

	r.Reset()
}

func TestUseOperatorNameResetToDefault(t *testing.T) {
	r := New()
	r.UseOperatorName(true)
	r.Start()
	done := r.Record("op_name", "11")
	done(0)
	r.Stop()

	rep := r.Report()
	if assert.Len(t, rep.Operators, 1) {
		assert.Equal(t, "op_name", rep.Operators[0].Name)
	}
	r.Reset()

	r2 := New()
	r2.Start()
	done2 := r2.Record("op_name", "11")
	done2(0)
	r2.Stop()

	rep2 := r2.Report()
	if assert.Len(t, rep2.Operators, 1) {
		assert.Equal(t, "11", rep2.Operators[0].Name)
	}
	r2.Reset()
}

func TestStartOffComputedAtRecordStart(t *testing.T) {
	r := New()
	r.Start()

	time.Sleep(20 * time.Millisecond)
	done := r.Record("op_name", "11")
	time.Sleep(20 * time.Millisecond)
	done(0)
	r.Stop()

	rep := r.Report()
	if assert.Len(t, rep.Operators, 1) {
		op := rep.Operators[0]
		// startOff 记录算子开始偏移，应接近 before-sleep 而不是 done 时刻。
		assert.Less(t, op.StartOff+op.Cost, rep.TotalCost+15*time.Millisecond)
	}
	r.Reset()
}

func TestReportReturnsCopy(t *testing.T) {
	r := New()
	r.Start()
	done := r.Record("op_name", "11")
	done(0)
	r.Stop()

	rep1 := r.Report()
	assert.Len(t, rep1.Operators, 1)
	rep1.Operators[0].Name = "changed_outside"

	rep2 := r.Report()
	if assert.Len(t, rep2.Operators, 1) {
		assert.Equal(t, "11", rep2.Operators[0].Name)
	}
	r.Reset()
}
