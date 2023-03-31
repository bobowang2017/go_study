package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, _ := context.WithTimeout(context.Background(), time.Minute*1)
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("协程退出, 停止了。..")
				return
			default:
				fmt.Println("协程运行中。..")
				time.Sleep(2 * time.Second)
			}
		}
	}(ctx)
	time.Sleep(time.Second * 70)
	fmt.Println("演示结束")
}
