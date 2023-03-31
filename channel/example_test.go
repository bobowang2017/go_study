package channel

import (
	"fmt"
	"strconv"
	"sync"
	"testing"
	"time"
)

func TestManyReceiver(t *testing.T) {
	total := 1000
	dataCh := make(chan string, 20)
	go func() {
		for i := 0; i < total; i++ {
			dataCh <- "task--" + strconv.Itoa(i)
		}
		close(dataCh)
	}()
	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			for {
				select {
				case name, ok := <-dataCh:
					if !ok {
						return
					}
					fmt.Println(fmt.Sprintf("线程%d接收数据:%s", idx, name))
					time.Sleep(3 * time.Second)
				}
			}
		}(i)
	}
	wg.Wait()
}
