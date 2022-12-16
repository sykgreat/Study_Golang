package main

import "fmt"

func main() {
	var a string
	// pair<static type: string,value: "fyy">
	a = "fyy"

	var allType interface{}
	// pair<type: string,value: "fyy">
	allType = a

	v, _ := allType.(string)
	fmt.Println(v)
}
