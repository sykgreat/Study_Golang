package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

func main() {
	// 泛型slice变量
	fmt.Println("======泛型slice变量======")
	type slice1[T int | string] []T
	var MySlice1 slice1[int] = []int{1, 2, 3} // 实例化
	MySlice2 := slice1[string]{"f", "y", "y"} // 实例化
	fmt.Println(MySlice1)
	fmt.Println(MySlice2)

	// 泛型map变量
	fmt.Println("======泛型map变量======")
	type map1[key int | string, value int | string] map[key]value
	var MyMap1 map1[int, int] = map[int]int{1: 1, 2: 2} // 实例化
	MyMap2 := map1[string, string]{"a": "b", "c": "d"}  // 实例化
	fmt.Println(MyMap1)
	fmt.Println(MyMap2)

	// 泛型struct变量（不支持匿名）
	type struct1[T int | string] struct {
		title string
		data  T
	}
	fmt.Println("======泛型struct变量======")
	var MyStruct1 struct1[int]
	MyStruct1.title = "bzd"
	MyStruct1.data = 1
	MyStruct2 := struct1[string]{"wsm", "bzd"}
	fmt.Println(MyStruct1) // 实例化
	fmt.Println(MyStruct2) // 实例化

	// 泛型变量嵌套
	fmt.Println("======泛型变量嵌套======")
	type struct2[S int | string, P map[S]string] struct {
		Name    string
		Content S
		Job     P
	}
	var MyStruct3 = struct2[int, map[int]string]{
		Name:    "small",
		Content: 1,
		Job:     map[int]string{1: "ss"},
	}
	MyStruct4 := struct2[string, map[string]string]{
		Name:    "big",
		Content: "bzd",
		Job:     map[string]string{"a": "b"},
	}
	fmt.Println(MyStruct3)
	fmt.Println(MyStruct4)

	//2个泛型变量之间的嵌套
	fmt.Println("======2个泛型变量之间的嵌套======")
	type slice3[T int | string | float64] []T
	type struct3[P int | string, V slice3[P]] struct {
		Name  P
		Title V
	}
	var MyStruct5 = struct3[int, slice3[int]]{
		Name:  1,
		Title: []int{1, 2, 3},
	}
	MyStruct6 := struct3[string, slice3[string]]{
		Name:  "bzd",
		Title: []string{"a", "b", "c"},
	}
	fmt.Println(MyStruct5)
	fmt.Println(MyStruct6)

	//泛型函数的调用
	fmt.Println("======泛型函数的调用======")
	a := 1
	b := 2
	a1 := 1.11
	b1 := 2.22
	sum1 := generic(a, b)
	sum2 := generic(a1, b1)
	fmt.Println(sum1)
	fmt.Println(sum2)

	//自定义类型约束
	fmt.Println("======自定义类型约束======")
	Foreach1([]int{1, 2, 3, 4})

	// 可排序的参数
	fmt.Println("======可排序的参数======")
	fmt.Println(Max[int](1, 2))

	// 比较==或者!==
	fmt.Println("======比较相不相等======")
	TF(1, 2)

	// 约束类型
	type MyInt interface{ ~int | ~int64 } // 这个约束的范围，不仅仅是int和int64本身，也包含只要最底层的是这2种类型的，都包含
}

func generic[T int | float64](a, b T) T {
	return a + b
}

type myInt interface {
	int | int8 | int16 | int32 | int64
}

type myUint interface {
	uint | uint8 | uint16 | uint32
}

type myFloat interface {
	float32 | float64
}

type myNumber interface {
	myInt | myUint | myFloat | string
}

// Foreach1 自定义类型参数
func Foreach1[T myNumber](list []T) {
	for _, t := range list {
		fmt.Println(t)
	}
}

// Foreach2 任何参数 any=interface{}
func Foreach2[T any](list []T) {
	for _, t := range list {
		fmt.Println(t)
	}
}

// Max 可排序的参数（满足这几个：(<、<=、>=、>)）
func Max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	} else {
		return b
	}
}

// TF 比较==或者!==
func TF[T comparable](a, b T) {
	if a == b {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}

// MyStruct 泛型指针结构体
type MyStruct[T interface{ *int | *float64 }] struct {
	Name T
}
