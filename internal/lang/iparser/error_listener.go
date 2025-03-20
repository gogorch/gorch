package iparser

import (
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

type ErrorListener struct {
	writer        io.Writer
	file          string
	errorAppender ErrorAppender
}

type ErrorAppender interface {
	AddError(e error)
}

func NewErrorListener(w io.Writer, fileName string, ea ErrorAppender) *ErrorListener {
	return &ErrorListener{
		writer:        w,
		file:          fileName,
		errorAppender: ea,
	}
}

// SyntaxError prints messages to System.err containing the
// values of line, charPositionInLine, and msg using
// the following format:
//
//	line <line>:<charPositionInLine> <msg>
func (el *ErrorListener) SyntaxError(
	recognizer antlr.Recognizer, offendingSymbol interface{},
	line, column int, msg string, e antlr.RecognitionException) {

	errorMsg := ""
	if len(el.file) > 0 {
		errorMsg = "file:" + el.file + ", "
	}
	errorMsg += "line " + strconv.Itoa(line) + ":" + strconv.Itoa(column) + " " + msg

	el.errorAppender.AddError(errors.New(errorMsg))
	if ct, ok := offendingSymbol.(*antlr.CommonToken); ok {
		el.underlineError(recognizer, ct, line, column)
	}
}

func (el *ErrorListener) underlineError(
	recognizer antlr.Recognizer, token *antlr.CommonToken,
	line, column int) {
	baseParser := recognizer.(*antlr.BaseParser)
	stream := baseParser.GetTokenStream()
	input := stream.GetTokenSource().GetInputStream().(fmt.Stringer).String()
	lines := strings.Split(input, "\n")
	preErrorLine := lines[line-2]
	if len(el.file) > 0 {
		fmt.Fprint(el.writer, el.file)
		fmt.Fprintln(el.writer, ":")
	}
	fmt.Fprintln(el.writer, line-1, " ", preErrorLine)
	errorLine := lines[line-1]
	fmt.Fprintln(el.writer, line, " ", errorLine)
	for i := 0; i < column; i++ {
		fmt.Fprint(el.writer, " ")
	}
	fmt.Fprint(el.writer, "     ")
	start := token.GetStart()
	stop := token.GetStop()
	if start >= 0 && stop >= 0 {
		for i := start; i <= stop; i++ {
			fmt.Fprintln(el.writer, "^")
		}
	}
	fmt.Fprintln(el.writer)
}

func (el *ErrorListener) ReportAmbiguity(
	recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool,
	ambigAlts *antlr.BitSet, configs *antlr.ATNConfigSet) {
	// fmt.Println("ReportAmbiguity")
}
func (el *ErrorListener) ReportAttemptingFullContext(
	recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int,
	conflictingAlts *antlr.BitSet, configs *antlr.ATNConfigSet) {
	// fmt.Println("ReportAttemptingFullContext")

}
func (el *ErrorListener) ReportContextSensitivity(
	recognizer antlr.Parser, dfa *antlr.DFA, startIndex int,
	stopIndex, prediction int, configs *antlr.ATNConfigSet) {
	// fmt.Println("ReportContextSensitivity")
}
