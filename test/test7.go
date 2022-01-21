package main

import (
	"fmt"
	"strconv"
)

func main() {
	var data = make(map[string]int)
	for i := 0; i < 100; i++ {
		data[strconv.Itoa(i)+"key"] = i
	}
	for k, v := range data {
		fmt.Println(v)
		delete(data, k)
	}
}
