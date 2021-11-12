package tool

import (
	"os"
	"os/signal"
	"syscall"
)

func QuitSignal(quitFunc func()) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	for s := range c {
		switch s {
		case syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
			quitFunc()
			return
		default:
			return
		}
	}
}
