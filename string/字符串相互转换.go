package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	str := "true"
	if res, err := strconv.ParseInt(str, 10, 32); err == nil {
		fmt.Println(res)
	} else {
		fmt.Println("error")
	}
	i := "10"
	if r, err := strconv.Atoi(i); err == nil {
		fmt.Println(reflect.TypeOf(r))
	} else {
		fmt.Println("exception")
	}

}
