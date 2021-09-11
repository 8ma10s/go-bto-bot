package main

import (
	"context"
	"errors"
	"github.com/ssym0614/go-bto-bot/internal/app/registry"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	reg := registry.NewRegistry()
	worker := reg.Worker()

	go func() {
		if err := worker.Start(ctx); err != nil && errors.Is(err, context.Canceled) {
			log.Println(err)
		}
	}()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	cancel()
	time.Sleep(5 * time.Second)
}
