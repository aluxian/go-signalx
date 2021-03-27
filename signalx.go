package signalx

import (
	"os"
	"os/signal"
	"syscall"
)

func WaitQuit() os.Signal {
	signalC := make(chan os.Signal, 1)
	signal.Notify(signalC, syscall.SIGINT, syscall.SIGTERM)
	s := <-signalC
	return s
}
