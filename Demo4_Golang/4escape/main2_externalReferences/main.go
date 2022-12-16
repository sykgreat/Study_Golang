package main

// 变量a在函数外部存在引用。
// 我们来分析一下执行过程：当函数执行完毕，对应的栈帧就被销毁，但是引用已经被返回到函数之外。如果这时外部通过引用地址取值，虽然地址还在，但是这块内存已经被释放回收了，这就是非法内存。
// 为了避免上述非法内存的情况，在这种情况下变量的内存分配必须分配到堆上。
func main() {
	_ = test()
}

func test() *int {
	a := 10
	return &a
}
