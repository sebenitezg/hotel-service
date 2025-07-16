package logger

import (
	"os"
	"sync"

	"go.elastic.co/ecszap"

	"go.uber.org/zap"
)

var lock = &sync.Mutex{}
var logger *zap.SugaredLogger

func NewLogger(debugMode bool) {
	lock.Lock()
	defer lock.Unlock()

	var logLevel = zap.InfoLevel

	if debugMode {
		logLevel = zap.DebugLevel
	}

	encoderConfig := ecszap.NewDefaultEncoderConfig()
	core := ecszap.NewCore(encoderConfig, os.Stdout, logLevel)
	theInstance := zap.New(core, zap.AddCaller())
	logger = theInstance.Sugar()
}

// GetLogger Returns the logger instance
func GetLogger(debugMode ...bool) *zap.SugaredLogger {
	var mode bool
	if len(debugMode) != 0 {
		mode = debugMode[0]
	}

	if logger == nil {
		NewLogger(mode)
	}

	return logger
}

// CloseLogger Flushes any pending logs
func CloseLogger() {
	_ = logger.Sync()
}
