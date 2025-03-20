package mlog

import (
	"bytes"
	"fmt"
	"log/slog"
)

type printLogger struct{}

func NewPrintLogger() Logger {
	return &printLogger{}
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

func (*printLogger) Debug(msg string, attrs ...LogValue) {
	slog.Debug(msg, logValueToAttr(attrs...)...)
}

func (*printLogger) AddInfo(attrs ...LogValue) {
	buf := &bytes.Buffer{}
	for i := range attrs {
		f := attrs[i]
		buf.WriteString(f.Key())
		buf.WriteString("[")
		fmt.Fprintf(buf, "%v", f.Value())
		buf.WriteString("] ")
	}
	slog.Info(buf.String())
}
func (*printLogger) Info(msg string, attrs ...LogValue) {
	slog.Info(msg, logValueToAttr(attrs...)...)
}

func (*printLogger) AddWarn(attrs ...LogValue) {
	buf := &bytes.Buffer{}
	for i := range attrs {
		f := attrs[i]
		buf.WriteString(f.Key())
		buf.WriteString("[")
		fmt.Fprintf(buf, "%v", f.Value())
		buf.WriteString("] ")
	}
	slog.Warn(buf.String())
}

func (*printLogger) Warn(msg string, attrs ...LogValue) {
	slog.Warn(msg, logValueToAttr(attrs...)...)
}

func (*printLogger) AddError(attrs ...LogValue) {
	buf := &bytes.Buffer{}
	for i := range attrs {
		f := attrs[i]
		buf.WriteString(f.Key())
		buf.WriteString("[")
		fmt.Fprintf(buf, "%v", f.Value())
		buf.WriteString("] ")
	}
	slog.Error(buf.String())
}

func (*printLogger) Error(msg string, attrs ...LogValue) {
	slog.Error(msg, logValueToAttr(attrs...)...)
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
