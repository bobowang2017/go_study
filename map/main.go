package main

import "fmt"

func main() {
	userInfo := map[string]interface{}{
		"name": "wangxiangbo",
		"age":  30,
		"sex":  true,
	}
	fmt.Println(userInfo["name"])
	fmt.Println(userInfo["age"])
	fmt.Println(userInfo["sex"])

	if bir, ok := userInfo["bir"]; ok {
		fmt.Println(bir)
	} else {
		fmt.Println("error")
	}
}
