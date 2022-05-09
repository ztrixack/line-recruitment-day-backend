package pkg

import (
	"context"
	"election-service/configs"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var startTime time.Time
var Context *context.Context
var cancelFunc *context.CancelFunc

func Init() {
	startTime = time.Now()
	ctx, cancel := context.WithCancel(context.Background())

	Context = &ctx
	cancelFunc = &cancel
}

func IsClosed() <-chan struct{} {
	return (*Context).Done()
}

func Started() {
	handleSigterm(func() {
		(*cancelFunc)()
		elapsed := time.Since(startTime)
		Info("OO", "Service stopped run_time=%v", elapsed)
	})

	elapsed := time.Since(startTime)
	Info("OO", "Service started setup_time=%v", elapsed)
	Info("OO", "Service listening at: %v", configs.Server.Port)
}

func Exit(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}

func handleSigterm(handleExit func()) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	go func() {
		<-c
		handleExit()
		os.Exit(1)
	}()
}
