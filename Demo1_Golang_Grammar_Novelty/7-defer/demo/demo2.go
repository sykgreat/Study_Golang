package demo

import "fmt"

// 知识点二：defer和return谁先谁后

func deferFunc() string {
	fmt.Println("defer func called...")
	return "fyy"
}

func returnFunc() string {
	fmt.Println("return func called...")
	return "fyy"
}

func DeferAndReturnFunc() string {
	defer deferFunc()
	return returnFunc()
}

//func main() {
//	DeferAndReturnFunc()
//}
//
// 输出：
// return func called...
// defer func called...
