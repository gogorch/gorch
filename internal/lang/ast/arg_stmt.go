package ast

import (
	"strconv"
	"strings"
	"time"

	"github.com/gorch/gorch/internal/lang/iantlr/alr"
)

type ConstantType int8

const (
	IntegerType ConstantType = iota
	StringType
	BooleanType
	DurationType
)

type ArgStmt struct {
	Type ConstantType

	Key string

	IntegerValue  []int64
	StringValue   []string
	BooleanValue  []bool
	DurationValue []time.Duration
}

type ArgStmtHolder interface {
	AcceptArgStmt(cons *ArgStmt) error
}

func (c *ArgStmt) AcceptIntegerLiteral(val int64) error {
	c.IntegerValue = append(c.IntegerValue, val)
	c.Type = IntegerType
	return nil
}

func (c *ArgStmt) AcceptStringLiteral(val string) error {
	c.StringValue = append(c.StringValue, val)
	c.Type = StringType
	return nil
}

func (c *ArgStmt) AcceptBoolLiteral(val bool) error {
	c.BooleanValue = append(c.BooleanValue, val)
	c.Type = BooleanType
	return nil
}

func (c *ArgStmt) AcceptDurationLiteral(val time.Duration) error {
	c.DurationValue = append(c.DurationValue, val)
	c.Type = DurationType
	return nil
}

func (c *ArgStmt) Description(blank string) string {
	s := new(strings.Builder)

	s.WriteString(c.Key)
	s.WriteString("=")

	switch c.Type {
	case IntegerType:
		if len(c.IntegerValue) > 1 {
			s.WriteString("[")
		}
		for idx, v := range c.IntegerValue {
			t := strconv.FormatInt(v, 10)
			s.WriteString(t)
			if idx < len(c.IntegerValue)-1 {
				s.WriteString(",")
			}
		}
		if len(c.IntegerValue) > 1 {
			s.WriteString("]")
		}
	case StringType:
		if len(c.StringValue) > 1 {
			s.WriteString("[")
		}
		for idx, v := range c.StringValue {
			s.WriteString("\"")
			s.WriteString(v)
			s.WriteString("\"")
			if idx < len(c.StringValue)-1 {
				s.WriteString(",")
			}
		}
		if len(c.StringValue) > 1 {
			s.WriteString("]")
		}
	case BooleanType:
		if len(c.BooleanValue) > 1 {
			s.WriteString("[")
		}
		for idx, v := range c.BooleanValue {
			if v {
				s.WriteString("true")
			} else {
				s.WriteString("false")
			}
			if idx < len(c.BooleanValue)-1 {
				s.WriteString(",")
			}
		}
		if len(c.BooleanValue) > 1 {
			s.WriteString("]")
		}
	case DurationType:
		for idx, v := range c.DurationValue {
			if v > 0 {
				s.WriteString(v.String())

				if idx < len(c.DurationValue)-1 {
					s.WriteString(",")
				}
			}
		}
		if len(c.DurationValue) > 1 {
			s.WriteString("]")
		}
	}
	return s.String()
}

// EnterArgStmt is called when entering the arg production.
func (s *GorchlangParserListener) EnterArgStmt(ctx *alr.ArgStmtContext) {
	if len(s.ParseErrors) > 0 {
		return
	}

	cons := &ArgStmt{Key: ctx.GetKey().GetText()}
	s.Stack.Push(cons)
}

// ExitArgStmt is called when exiting the arg production.
func (s *GorchlangParserListener) ExitArgStmt(ctx *alr.ArgStmtContext) {
	if len(s.ParseErrors) > 0 {
		return
	}
	cons := s.Stack.Pop().(*ArgStmt)
	holder := s.Stack.Peek().(ArgStmtHolder)
	err := holder.AcceptArgStmt(cons)
	if err != nil {
		s.AddError(err)
	}
}

type StringLiteralHolder interface {
	AcceptStringLiteral(val string) error
}

// EnterStringLiteral is called when production stringLiteral is entered.
func (s *GorchlangParserListener) EnterStringLiteral(ctx *alr.StringLiteralContext) {}

// ExitStringLiteral is called when production stringLiteral is exited.
func (s *GorchlangParserListener) ExitStringLiteral(ctx *alr.StringLiteralContext) {
	if len(s.ParseErrors) > 0 {
		return
	}
	sval := ctx.GetText()
	if len(sval) > 2 {
		sval = sval[1 : len(sval)-1]
	} else {
		sval = ""
	}

	var err error
	holder := s.Stack.Peek().(StringLiteralHolder)
	err = holder.AcceptStringLiteral(sval)
	if err != nil {
		s.AddError(err)
	}
}

type BoolLiteralHolder interface {
	AcceptBoolLiteral(val bool) error
}

// EnterBoolLiteral is called when production booleanLiteral is entered.
func (s *GorchlangParserListener) EnterBoolLiteral(ctx *alr.BoolLiteralContext) {}

// ExitBoolLiteral is called when production booleanLiteral is exited.
func (s *GorchlangParserListener) ExitBoolLiteral(ctx *alr.BoolLiteralContext) {
	if len(s.ParseErrors) > 0 {
		return
	}
	holder := s.Stack.Peek().(BoolLiteralHolder)
	sval := strings.ToLower(ctx.GetText())
	bval := false
	if sval == "true" {
		bval = true
	}
	err := holder.AcceptBoolLiteral(bval)
	if err != nil {
		s.AddError(err)
	}
}

type IntegerLiteralHolder interface {
	AcceptIntegerLiteral(val int64) error
}

// EnterInteger is called when production integer is entered.
func (s *GorchlangParserListener) EnterIntegerLiteral(ctx *alr.IntegerLiteralContext) {}

// ExitInteger is called when production integer is exited.
func (s *GorchlangParserListener) ExitIntegerLiteral(ctx *alr.IntegerLiteralContext) {
	if len(s.ParseErrors) > 0 {
		return
	}
	val, err := strconv.ParseInt(ctx.GetText(), 10, 64)
	if err != nil {
		s.AddError(err)
		return
	}
	holder := s.Stack.Peek().(IntegerLiteralHolder)
	err = holder.AcceptIntegerLiteral(val)
	if err != nil {
		s.AddError(err)
	}
}

type DurationLiteralHolder interface {
	AcceptDurationLiteral(val time.Duration) error
}

// EnterDuration is called when production integer is entered.
func (s *GorchlangParserListener) EnterDurationLiteral(ctx *alr.DurationLiteralContext) {}

// ExitDuration is called when production integer is exited.
func (s *GorchlangParserListener) ExitDurationLiteral(ctx *alr.DurationLiteralContext) {
	if len(s.ParseErrors) > 0 {
		return
	}
	val, err := time.ParseDuration(ctx.GetText())
	if err != nil {
		s.AddError(err)
		return
	}
	holder := s.Stack.Peek().(DurationLiteralHolder)
	err = holder.AcceptDurationLiteral(val)
	if err != nil {
		s.AddError(err)
	}
}
