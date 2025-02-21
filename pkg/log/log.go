package log

import "time"

type Logger interface {
	Info(msg string, fn ...LoggerContext)
	Error(msg string, fn ...LoggerContext)
	Warn(msg string, fn ...LoggerContext)
	Debug(msg string, fn ...LoggerContext)
	Fatal(msg string, fn ...LoggerContext)
	Panic(msg string, fn ...LoggerContext)
}

type LoggerContextFn func(LoggerContext)

type LoggerContext interface {
	Any(key string, value any)
	Bool(key string, value bool)
	Bytes(key string, value []byte)
	String(key string, value string)
	Float64(key string, value float64)
	Int64(key string, value int64)
	Uint64(key string, value uint64)
	Time(key string, value time.Time)
	Duration(key string, value time.Duration)
	Error(key string, err error)
}

type Loggable interface {
	AsLog() any
}
