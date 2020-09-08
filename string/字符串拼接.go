package main

import (
	"bytes"
	"fmt"
	"strings"
)

var s string

func main() {
	/**
	1、直接使用+拼接
	*/
	a := "hello"
	b := "world"
	res := ""
	for i := 0; i < 10; i++ {
		res += a + b
	}
	fmt.Println(s)
	s = res
	fmt.Println(s)
	/**
	2、fmt.Sprintf()
	*/
	res = fmt.Sprintf("%s%s%d", a, b, 1)
	fmt.Println(res)
	/**
	3. strings.Join()
	*/
	strs := []string{"hello", "world", "您好"}
	fmt.Println(strings.Join(strs, "-"))
	/**
	4. buffer.WriteString()
	*/
	var buffer bytes.Buffer
	buffer.WriteString("hello")
	buffer.WriteString(",")
	buffer.WriteString("world")
	res = buffer.String()
	fmt.Println(res)
}

/**
参考网址：https://www.jianshu.com/p/df92c0ee6cc8
*/
