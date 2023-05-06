package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)
	done := make(chan bool)
	for x := 0; x < 5; x++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-done:
					fmt.Println("Goroutine exiting")
					return
				case val := <-ch:
					fmt.Printf("Received value: %d\n", val)
				default: // Non-blocking channel operation
				}
			}
		}()
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			ch <- i
		}
		close(done)
	}()
	wg.Wait()
	fmt.Println("All goroutines have completed")
}
