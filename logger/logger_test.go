package logger

import (
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	logger := New(WithLogFormatter("json"))
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()
	for range ticker.C {
		logger.Infof("time:%d", time.Now().Unix())
	}
}
