package ast

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
)

type Pos struct {
	File        string
	StartLine   int
	StartColunm int
	StopLine    int
	StopColunm  int
}

func ToPos(ap interface {
	GetStart() antlr.Token
	GetStop() antlr.Token
}, file string) Pos {
	start := ap.GetStart()
	stop := ap.GetStop()
	return Pos{
		File:        file,
		StartLine:   start.GetLine(),
		StartColunm: start.GetColumn(),
		StopLine:    stop.GetLine(),
		StopColunm:  stop.GetColumn(),
	}
}

func (p Pos) String() string {
	return fmt.Sprintf("%s %d", p.File, p.StartLine)
}
