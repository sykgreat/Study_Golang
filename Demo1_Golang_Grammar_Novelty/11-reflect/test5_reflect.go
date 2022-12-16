package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (u User) Call() {
	fmt.Println("user is called ...")
	fmt.Printf("%v\n", u)
}

func main() {
	user := User{Id: 1, Name: "fyy", Age: 18}
	DoFiledAndMethod(user)
}

func DoFiledAndMethod(input interface{}) {
	// 获取input的type
	inputType := reflect.TypeOf(input)
	fmt.Printf("input type: %v\n", inputType)
	fmt.Printf("input type: %v\n", inputType.Name())
	fmt.Println()

	// 获取input的value
	inputValue := reflect.ValueOf(input)
	fmt.Printf("input value: %v\n", inputValue)
	fmt.Println()

	// 通过type获取里面的字段
	// 1. 获取interface的reflect.Type，通过Type得到NumField，进行遍历
	// 2. 得到每一个field，数据类型
	// 3. 通过field有一个Interface()方法得到对应的value
	for i := 0; i < inputValue.NumField(); i++ {
		field := inputType.Field(i)
		value := inputValue.Field(i).Interface()

		fmt.Printf("%s: %v = %v\n", field.Name, field.Type, value)
	}
	fmt.Println()

	// 通过type获取里面的方法并调用
	for i := 0; i < inputType.NumMethod(); i++ {
		m := inputType.Method(i)
		fmt.Printf("%s: %v\n", m.Name, m.Type)
	}
}
