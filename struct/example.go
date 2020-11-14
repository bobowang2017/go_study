package main

import (
	"fmt"
	"reflect"
)

type UserInfo struct {
	ID   int    `json:"id" form:"id"`
	Name string `json:"name" form:"name"`
	Age  int    `json:"age" form:"age"`
	Sex  bool   `json:"sex" form:"sex"`
}

func (u UserInfo) getUserName(s string) string {
	fmt.Println(s)
	u.Name = "world"
	return u.Name
}

func main() {
	u := UserInfo{
		123,
		"wangxiangbo",
		23,
		true,
	}
	//v := reflect.ValueOf(&u)
	//t := v.Type()
	//kind := t.Kind()
	//fmt.Println(kind)
	//for i := 0; i < v.NumField(); i++ {
	//	field := v.Field(i)
	//	//打印字段的名称、类型以及值
	//	fmt.Printf("name:%s type:%v value:%v\n",
	//		t.Field(i).Name, field.Type().Kind(), field.Interface())
	//
	//}
	//tag :=t.Elem().Field(0)
	//fmt.Printf("tag json=%s\n", tag.Tag.Get("json"))
	s := reflect.ValueOf(&u).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i,
			typeOfT.Field(i).Name, f.Type(), f.Interface())
	}
}
