package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 1; i < 10000; i++ {
		go func(i int) {
			fmt.Println("Hello World", i)
		}(i)
	}
	time.Sleep(1)
	fmt.Println("Main")
}
