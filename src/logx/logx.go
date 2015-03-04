package logx

import (
	"log"
	"os"
)

const (
	LogFilePath = "log/running.log"
	InfoLevel   = 1 << iota //Info level. log Info/Warning/Error
	WarningLevel
	ErrorLevel
)

var (
	red    = string([]byte{27, 91, 57, 49, 109})
	green  = string([]byte{27, 91, 57, 50, 109})
	yellow = string([]byte{27, 91, 57, 51, 109})
	reset  = string([]byte{27, 91, 48, 109})
)

type Logger struct {
	BuiltInLogger *log.Logger
	LogLevel      int
}

func New(logLevel int) *Logger {
	builtInLogger := log.New(os.Stdout, "", log.LstdFlags)
	return &Logger{BuiltInLogger: builtInLogger, LogLevel: logLevel}
}

func (logger *Logger) Info(message interface{}) {
	if logger.LogLevel <= InfoLevel {
		logger.BuiltInLogger.SetFlags(log.LstdFlags)
		logger.BuiltInLogger.Println(green, "[INFO]", reset, message)
	}
}

func (logger *Logger) Warn(message interface{}) {
	if logger.LogLevel <= WarningLevel {
		logger.BuiltInLogger.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
		logger.BuiltInLogger.Println(yellow, "[WARN]", reset, message)
	}
}

func (logger *Logger) Error(message interface{}) {
	if logger.LogLevel <= ErrorLevel {
		logger.BuiltInLogger.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
		logger.BuiltInLogger.Println(red, "[ERROR]", reset, message)
	}
}
