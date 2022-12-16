package main

// 我们虽然在代码段中给变量 l 赋值了1，但是编译期间只能识别到初始化int类型切片时，传入的长度和容量是变量l，编译期并不能确定变量l的值，所以发生了逃逸，会把内存分配到堆中。
func main() {
	test()
}

func test() {
	l := 1
	a := make([]int, l, l)
	for i := 0; i < l; i++ {
		a = append(a, i)
	}
}
