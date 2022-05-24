package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGTERM)

	go func() {
		for {
			time.Sleep(3 * time.Second)
		}
	}()

	for {
		select {
		case <-ch:
			fmt.Println()
		}
	}
}
