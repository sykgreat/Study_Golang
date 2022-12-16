package main

type A struct {
	a int
}

func main() {
	a := make([]int, 1, 1)
	b := new(A)
	b.a = 1
	a = append(a, 1)
}
