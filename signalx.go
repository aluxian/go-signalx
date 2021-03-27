package signalx

import (
	"os"
	"os/signal"
	"syscall"
)

func Notifications(sig ...os.Signal) <-chan os.Signal {
	signalC := make(chan os.Signal, 1)
	signal.Notify(signalC, sig...)
	return signalC
}

func QuitC() <-chan os.Signal {
	return Notifications(syscall.SIGINT, syscall.SIGTERM)
}

func WaitQuit() os.Signal {
	s := <-QuitC()
	return s
}
