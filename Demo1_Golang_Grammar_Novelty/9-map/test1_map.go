package main

import (
	"fmt"
)

func main() {
	// ====> 第一种声明方式
	// 声明myMap1是一个map类型 key是string val是string
	var myMap1 map[string]string
	if myMap1 == nil {
		fmt.Println("myMap1 is nil")
	}
	// 在使用map前，需要先用make给map分配空间
	myMap1 = make(map[string]string, 10)

	myMap1["syk"] = "java"
	myMap1["sky"] = "golang"
	myMap1["hpq"] = "fyy"
	fmt.Println(myMap1)
	fmt.Println()

	// ====> 第二种声明方式
	myMap2 := make(map[string]string)
	myMap2["syk"] = "java"
	myMap2["sky"] = "golang"
	myMap2["hpq"] = "fyy"
	fmt.Println(myMap2)
	fmt.Println()

	// ====> 第三种声明方式
	myMap3 := map[string]string{
		"syk": "java",
		"sky": "golang",
		"hpq": "fyy",
	}
	fmt.Println(myMap3)
	fmt.Println()
}
