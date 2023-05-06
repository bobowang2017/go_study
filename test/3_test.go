package main

import (
	"context"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// 给定一个RPC模拟调用方法，耗时在5-15s内，您需要在函数中模拟调用rpc方法，设置超时时间10秒，未超时则输出结果，超时则输出timeout

func TestTimer(t *testing.T) {
	timer := time.NewTimer(10 * time.Second)
	defer timer.Stop()
	resCh := make(chan int, 1)
	go func() {
		resCh <- rpc()
	}()
	select {
	case <-timer.C:
		fmt.Println("timeout")
	case v, _ := <-resCh:
		fmt.Println(v)
	}
}

func TestContext(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	resCh := make(chan int, 1)
	go func() {
		resCh <- rpc()
	}()
	select {
	case <-ctx.Done():
		fmt.Println("timeout")
	case v, _ := <-resCh:
		fmt.Println(v)
	}
}

func TestChannel(t *testing.T) {
	resCh := make(chan int, 1)
	go func() {
		resCh <- rpc()
	}()
	for i := 0; i < 10; i++ {
		select {
		case v, _ := <-resCh:
			fmt.Println(v)
			return
		default:
			time.Sleep(1 * time.Second)
			fmt.Println("Sleep One Second")
		}
	}
	fmt.Println("timeout")
}

func rpc() int {
	x := RandomInt(8, 15)
	time.Sleep(time.Duration(x) * time.Second)
	fmt.Println(fmt.Sprintf("I Sleep %d s", x))
	return x
}

func RandomInt(min, max int) int {
	if min < 0 || max < 0 {
		return 0
	}
	if min > max {
		return 0
	}
	if min == max {
		return min
	}
	return rand.Int()%max + min
}
