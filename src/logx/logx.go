package logx

import (
	"fmt"
	"log"
	"os"
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
	std    = log.New(os.Stderr, "", log.LstdFlags)
)

func Info(message interface{}) {
	if LogLevel > InfoLevel {
		return
	}
	std.SetFlags(log.LstdFlags)
	std.Println(green, "[INFO]", reset, message)
}

func Warn(message interface{}) {
	if LogLevel > WarningLevel {
		return
	}
	std.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	std.Output(3, fmt.Sprintln(yellow, "[WARN]", reset, message))
}

func Error(message interface{}) {
	if LogLevel > ErrorLevel {
		return
	}
	std.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	std.Output(3, fmt.Sprintln(red, "[ERROR]", reset, message))
}
