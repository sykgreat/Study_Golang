package demo

import "fmt"

// 知识点一：defer的执行顺序

func S() {
	fmt.Println("s f")
}

func Y() {
	fmt.Println("y y")
}
func K() {
	fmt.Println("k y")
}

//func main() {
//	defer demo.S() 3
//	defer demo.Y() 2
//	defer demo.K() 1
//}

// 输出：
// k y
// y y
// s f
