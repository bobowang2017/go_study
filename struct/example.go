package main

import (
	"fmt"
	"github.com/google/uuid"
	"testing"
	"unsafe"
)

func TestBenchmarkSlice(t *testing.T) {
	dataMap := make(map[string]string, 100)
	cnt := 0
	for cnt < 1000000 {
		dataMap[uuid.New().String()] = uuid.New().String()
		cnt += 1
	}
	fmt.Println(unsafe.Sizeof(dataMap))
}
