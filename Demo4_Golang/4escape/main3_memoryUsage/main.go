package main

func main() {
	test1()
	test2()
}

// 占用内存过大 发生逃逸
func test1() {
	a := make([]int, 1000000)
	for i := 0; i < 1000000; i++ {
		a = append(a, i)
	}
}

// 不发生逃逸
func test2() {
	a := make([]int, 10)
	for i := 0; i < 10; i++ {
		a = append(a, i)
	}
}
