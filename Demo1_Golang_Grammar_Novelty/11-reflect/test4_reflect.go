package main

import (
	"fmt"
	"reflect"
)

func reflectNum(arg interface{}) {
	fmt.Println("type: ", reflect.TypeOf(arg))
	fmt.Println("value: ", reflect.ValueOf(arg))
}

func main() {
	var num float64 = 3.141592653589793
	reflectNum(num)
}
