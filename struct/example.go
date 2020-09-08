package main

import "fmt"

type UserInfo struct {
	ID   int
	Name string
	age  int
	sex  bool
}

func (u UserInfo) getUserName(s string) string {
	fmt.Println(s)
	u.Name = "world"
	return u.Name
}

func main() {
	var userInfo UserInfo
	userInfo.ID = 1
	userInfo.sex = false
	userInfo.Name = "hello"
	name := userInfo.getUserName("bobo")
	fmt.Println(name)
	fmt.Println(userInfo.Name)
}
