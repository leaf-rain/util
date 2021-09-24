package logger

import (
	"testing"
)

func TestGetLogger(t *testing.T) {
	log := GetLogger()
	log.Info("test")
}
