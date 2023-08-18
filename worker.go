package main

import (
	"context"
	"time"
)

func NewWorker(ctx context.Context, period time.Duration, f func()) {
	if period <= 0 {
		log.Println("Error: non-positive interval for NewTicker")
		return
	}
	tick := time.NewTicker(period)
	GetProgress() // Call GetProgress function and log the progress
	f()
	defer tick.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-tick.C:
			f()
		}
	}
}
