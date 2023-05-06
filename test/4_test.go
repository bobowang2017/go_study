package main

import (
	"fmt"
	"testing"
)

func TestFor(t *testing.T) {
	data := make(chan int, 1)
	for {
		select {
		case <-data:
			fmt.Println("success")
		default:
			fmt.Println("hello")
		}
	}
}
