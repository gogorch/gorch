package mlog

import (
	"reflect"
	"time"
)

type Logger interface {
	AddDebug(...LogValue)

	Debug(string, ...LogValue)

	AddInfo(...LogValue)

	Info(string, ...LogValue)

	AddWarn(...LogValue)

	Warn(string, ...LogValue)

	AddError(...LogValue)

	Error(string, ...LogValue)
}

// LogKind 表明当前 logValue 的类型
type LogKind uint8

const (
	KindUnknown LogKind = iota
	KindBytes
	KindBool
	KindBytesString
	KindDuration
	KindFloat64
	KindFloat32
	KindInt
	KindInt64
	KindInt32
	KindInt16
	KindInt8
	KindString
	KindTime
	KindUint
	KindUint64
	KindUint32
	KindUint16
	KindUint8
	KindUintptr
	KindReflect
	KindReflectType
	KindError
	KindAny
)

var kindStrings = []string{
	"Unknown",
	"Bytes",
	"Bool",
	"BytesString",
	"Duration",
	"Float64",
	"Float32",
	"Int",
	"Int64",
	"Int32",
	"Int16",
	"Int8",
	"String",
	"Time",
	"Uint",
	"Uint64",
	"Uint32",
	"Uint16",
	"Uint8",
	"Reflect",
	"Error",
}

func (k LogKind) String() string {
	if int(k) < len(kindStrings) {
		return kindStrings[k]
	}
	return "<unknown mlog.LogKind>"
}

type LogValue interface {
	Key() string
	Value() any
}

type logValue struct {
	key string
	val any
	typ LogKind
}

func (f *logValue) Key() string {
	return f.key
}

func (f *logValue) Value() any {
	return f.val
}

func Bytes(key string, value []byte) LogValue           { return &logValue{key, value, KindBytes} }
func Bool(key string, value bool) LogValue              { return &logValue{key, value, KindBool} }
func BytesString(key string, value []byte) LogValue     { return &logValue{key, value, KindBytesString} }
func Duration(key string, value time.Duration) LogValue { return &logValue{key, value, KindDuration} }
func Float64(key string, value float64) LogValue        { return &logValue{key, value, KindFloat64} }
func Float32(key string, value float32) LogValue        { return &logValue{key, value, KindFloat32} }
func Int(key string, value int) LogValue                { return &logValue{key, value, KindInt} }
func Int64(key string, value int64) LogValue            { return &logValue{key, value, KindInt64} }
func Int32(key string, value int32) LogValue            { return &logValue{key, value, KindInt32} }
func Int16(key string, value int16) LogValue            { return &logValue{key, value, KindInt16} }
func Int8(key string, value int8) LogValue              { return &logValue{key, value, KindInt8} }
func String(key string, value string) LogValue          { return &logValue{key, value, KindString} }
func Time(key string, value time.Time) LogValue         { return &logValue{key, value, KindTime} }
func Uint(key string, value uint) LogValue              { return &logValue{key, value, KindUint} }
func Uint64(key string, value uint64) LogValue          { return &logValue{key, value, KindUint64} }
func Uint32(key string, value uint32) LogValue          { return &logValue{key, value, KindUint32} }
func Uint16(key string, value uint16) LogValue          { return &logValue{key, value, KindUint16} }
func Uint8(key string, value uint8) LogValue            { return &logValue{key, value, KindUint8} }
func Reflect(key string, value reflect.Value) LogValue  { return &logValue{key, value, KindReflect} }
func ReflectType(key string, value reflect.Type) LogValue {
	return &logValue{key, value, KindReflectType}
}
func Error(key string, value error) LogValue { return &logValue{key, value, KindError} }
func Any(key string, value any) LogValue     { return &logValue{key, value, KindAny} }
