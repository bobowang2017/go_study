package main

import (
	"fmt"
	"strconv"
	"sync"
	"testing"
	"time"
)

func TestJenkins(t *testing.T) {
	total := 1000
	dataCh := make(chan string, 20)
	go func() {
		for i := 0; i < total; i++ {
			dataCh <- "task--" + strconv.Itoa(i)
		}
		fmt.Println("关闭通道")
		close(dataCh)
	}()
	time.Sleep(3 * time.Second)
	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			for {
				select {
				case name, ok := <-dataCh:
					if !ok {
						fmt.Println("接收关闭通道信号")
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

func DelJenkinsUser(username string) error {
	//var (
	//	resp *HttpRequest.Response
	//	err  error
	//)
	////cookies := map[string]string{
	////	"JSESSIONID.fc4ee955": "node0bbo4dhqnwsjk1pmj8k70ilfpx72756.node0",
	////	"JSESSIONID.5ffa61e0": "node06wai8bhlexwo1t6iz1m2dg76e1.node0",
	////	"JSESSIONID.ac32ccf9": "node0160zv1hnwh3ji14u7uo83vlg6h2.node0",
	////	"JSESSIONID.3e47cd30": "node03l7a172dxoc3lp2bz70sz81o1.node0",
	////	"JSESSIONID.c27178a9": "node016jtjd6sfmvzpc97ippbdbeju4.node0",
	////	"JSESSIONID.31d6b364": "node01bc5p0slgyl401bhjiomq24bhz2.node0",
	////	"JSESSIONID.b6324199": "node0zpvxxes4l6n4rkgyaz7uhqus2.node0",
	////	"JSESSIONID.edc0c727": "node01vcgu3yxnnel210e1lmo9hmvjw3.node0",
	////	"screenResolution":    "1536x864",
	////}
	//cookies := map[string]string{
	//	"JSESSIONID.fc4ee955": "node012pn880ppzddbfl3elqg2exs766492.node0",
	//	"sidebarStatus":       "0",
	//	"JSESSIONID.edc0c727": "node0xa26gkuqvfuf15yrjbgklkc5r126.node0",
	//	"screenResolution":    "1920x1080",
	//}
	//req := HttpRequest.NewRequest().SetTimeout(3).SetCookies(cookies)
	//url := fmt.Sprintf("https://devops.cm-iov.com/jenkins/user/%s/doDelete", username)
	//params := fmt.Sprintf(`{"Jenkins-Crumb": "533568540156f1bb4ad6a7f3289a721c9af38f647b59d00fa1d9d2cf2db4b1c3"}`)
	//if resp, err = req.JSON().Post(url, params); err != nil {
	//	return err
	//}
	//fmt.Println(resp.StatusCode())
	return nil
}

func TestHello(t *testing.T) {
	tasks := []string{"taskA", "taskB", "taskC"}
	resCh := make(chan string, 3)
	var wg sync.WaitGroup
	wg.Add(len(tasks))

	for i := 0; i < len(tasks); i++ {
		go func(idx int, name string) {
			defer wg.Done()
			fmt.Printf("线程%d处理任务%s\n", idx, name)
			time.Sleep(1 * time.Second)
			resCh <- fmt.Sprintf("任务%s处理完毕", name)
		}(i, tasks[i])
	}

	wg.Wait()
	close(resCh)
	for data := range resCh {
		fmt.Println(data)
	}
}

func TestName(t *testing.T) {
	var wg sync.WaitGroup
	ch := make(chan int, 10)
	done := make(chan bool)
	for x := 0; x < 5; x++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-done:
					fmt.Println("Goroutine exiting")
					return
				case val := <-ch:
					fmt.Printf("Received value: %d\n", val)
					//default: // Non-blocking channel operation
				}
			}
		}()
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			ch <- i
		}
		close(done)
	}()
	wg.Wait()
	fmt.Println("All goroutines have completed")
}
