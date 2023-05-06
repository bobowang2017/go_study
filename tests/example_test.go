package tests

import (
	"github.com/google/uuid"
	"testing"
)

func BenchmarkSlice(b *testing.B) {
	dataMap := make(map[string]string, 100)
	cnt := 0
	for cnt < 100000 {
		dataMap[uuid.New().String()] = uuid.New().String()
		cnt += 1
	}
	//fmt.Println(unsafe.Sizeof(dataMap))
}
