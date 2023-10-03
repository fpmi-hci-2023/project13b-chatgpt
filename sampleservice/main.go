package main

import (
	"context"
	"log/slog"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	slog.Info("Hello world")

	timer := time.NewTicker(10 * time.Second)

	go func() {
		for {
			select {
			case <-timer.C:
				slog.Info("Tick")
			case <-ctx.Done():
				slog.Info("Context done")
				return
			}
		}
	}()

	<-ctx.Done()
	slog.Info("exiting now")
}
