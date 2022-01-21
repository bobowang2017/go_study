package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name  string
	Sex   int
	Age   int
	Score float32
}

func ReflectTypeOf(a interface{}) {
	t := reflect.TypeOf(a)
	fmt.Printf("type of a is:%v\n", t)
	k := t.Kind()
	switch k {
	case reflect.Int:
		fmt.Printf("a is int\n")
	case reflect.String:
		fmt.Printf("a is string\n")
	}
}

func ReflectValueOf(a interface{}) {
	v := reflect.ValueOf(a)
	t := v.Type()
	k := t.Kind()
	switch k {
	case reflect.Int64:
		fmt.Printf("a is Int64, store value is:%d\n", v.Int())
	case reflect.String:
		fmt.Printf("a is String, store value is:%s\n", v.String())
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			field := v.Field(i)
			//打印字段的名称、类型以及值
			fmt.Printf("name:%s type:%v value:%v\n",
				t.Field(i).Name, field.Type().Kind(), field.Interface())
		}
	}
}

func main() {
	//ReflectTypeOf(5)
	var s Student = Student{
		Name:  "BigOrange",
		Sex:   1,
		Age:   10,
		Score: 80.1,
	}
	ReflectValueOf(s)
}
