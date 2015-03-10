package logx

import (
	"log"
)

const LogLevel = 0 // current Log Level

const (
	InfoLevel = 1 << iota //Info level. log Info/Warning/Error
	WarningLevel
	ErrorLevel
)

var (
	red    = string([]byte{27, 91, 57, 49, 109})
	green  = string([]byte{27, 91, 57, 50, 109})
	yellow = string([]byte{27, 91, 57, 51, 109})
	reset  = string([]byte{27, 91, 48, 109})
)

func Info(message interface{}) {
	if LogLevel > InfoLevel {
		return
	}
	log.SetFlags(log.LstdFlags)
	log.Println(green, "[INFO]", reset, message)
}

func Warn(message interface{}) {
	if LogLevel > WarningLevel {
		return
	}
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	log.Println(yellow, "[WARN]", reset, message)
}

func Error(message interface{}) {
	if LogLevel > ErrorLevel {
		return
	}
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	log.Println(red, "[ERROR]", reset, message)
}
