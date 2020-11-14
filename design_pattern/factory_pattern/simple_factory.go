package main

import (
	"fmt"
)

type Operation interface {
	Exe(int, int) int
}

type OperationAdd struct{}
type OperationSub struct{}
type OperationMul struct{}
type OperationDiv struct{}

func (o *OperationAdd) Exe(a int, b int) int {
	return a + b
}

func (o *OperationSub) Exe(a int, b int) int {
	return a - b
}

func (o *OperationMul) Exe(a int, b int) int {
	return a * b
}

func (o *OperationDiv) Exe(a int, b int) int {
	return a / b
}

func OperationFactory(oper string) Operation {
	switch oper {
	case "+":
		return &OperationAdd{}
	case "-":
		return &OperationSub{}
	case "*":
		return &OperationMul{}
	case "/":
		return &OperationDiv{}
	}

	return nil
}

func main() {
	fmt.Println(
		OperationFactory("+").Exe(10, 5),
		OperationFactory("-").Exe(10, 5),
		OperationFactory("*").Exe(10, 5),
		OperationFactory("/").Exe(10, 5))
}
