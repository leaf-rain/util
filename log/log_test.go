package log

import (
	"testing"
)

func TestGetLogger(t *testing.T) {
	ReloadConfig(&Options{
		LogFileDir: "",
		AppName:    "test",
		Platform:   "test",
		MaxSize:    0,
		MaxBackups: 0,
		MaxAge:     0,
		Level:      "info",
		ToFile:     true,
	})
	logger.Info("test")
}
