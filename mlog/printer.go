package mlog

import (
	"bytes"
	"fmt"
	"log/slog"
	"os"
)

type printLogger struct {
	l *slog.Logger
}

func NewPrintLogger() Logger {
	var lvl slog.LevelVar
	lvl.Set(slog.LevelDebug)
	opts := slog.HandlerOptions{
		Level: &lvl,
	}
	l := slog.New(slog.NewJSONHandler(os.Stderr, &opts))

	return &printLogger{l}
}

func (*printLogger) AddDebug(attrs ...LogValue) {
	buf := &bytes.Buffer{}
	for i := range attrs {
		f := attrs[i]
		buf.WriteString(f.Key())
		buf.WriteString("[")
		fmt.Fprintf(buf, "%v", f.Value())
		buf.WriteString("] ")
	}
	slog.Debug(buf.String())
}

func (pl *printLogger) Debug(msg string, attrs ...LogValue) {
	pl.l.Debug(msg, logValueToAttr(attrs...)...)
}

func (pl *printLogger) Info(msg string, attrs ...LogValue) {
	pl.l.Info(msg, logValueToAttr(attrs...)...)
}

func (pl *printLogger) Warn(msg string, attrs ...LogValue) {
	pl.l.Warn(msg, logValueToAttr(attrs...)...)
}

func (pl *printLogger) Error(msg string, attrs ...LogValue) {
	pl.l.Error(msg, logValueToAttr(attrs...)...)
}

func logValueToAttr(attrs ...LogValue) []any {
	fields := []any{}
	for i := range attrs {
		f := attrs[i]
		fields = append(fields, slog.Any(f.Key(), f.Value()))
	}
	return fields
}

var _ Logger = (*printLogger)(nil)
