package main

import (
	"fmt"
)

var chan1 = make(chan bool)
var chan2 = make(chan bool)
var tag = make(chan bool)

func func1() {
	i := 0
	for {
		i++
		<-chan1
		fmt.Print(i)
		i++
		fmt.Print(i)
		chan2 <- true
	}
}

func func2() {
	for i := 'A'; i <= 'Z'; i += 2 {
		<-chan2
		fmt.Print(string(i))
		fmt.Print(string(i + 1))
		chan1 <- true
	}
	tag <- true
}

func main() {
	go func1()
	go func2()
	chan1 <- true
	<-tag
}
