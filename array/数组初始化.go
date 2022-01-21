package main

import (
	"fmt"
	"reflect"
)

func main() {
	month := [...]string{
		1:  "january",
		2:  "february",
		3:  "march",
		4:  "april",
		5:  "may",
		6:  "june",
		7:  "july",
		8:  "august",
		9:  "september",
		10: "october",
		11: "november",
		12: "December",
	}
	fmt.Println(month)
	fmt.Println(len(month))
	fmt.Println(reflect.TypeOf(month))
	fmt.Println(reflect.TypeOf(month[4:]))
}
