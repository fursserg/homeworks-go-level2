package main

// Здесь часть домашки с счетчиком до 1000

import (
	"fmt"
	"sync"
)

func main() {
	var count int
	wg := sync.WaitGroup{}

	for i := 0; i <= 100; i++ {
		wg.Add(1)
		go func() {
			for {
				if count >= 1000 {
					wg.Done()
					return
				}

				count++
			}
		}()
	}

	wg.Wait()

	fmt.Println(count)
}
