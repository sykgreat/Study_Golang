package main

import (
	"fmt"
)

type Human struct {
	Name string
	Sex  string
}

func (h *Human) Eat() {
	fmt.Println("Human.Eat()...")
}
func (h *Human) Sleep() {
	fmt.Println("Human.Sleep()...")
}

type SuperMan struct {
	Human // 继承Human类
	level int
}

// Eat 重定义父类方法
func (s *SuperMan) Eat() {
	fmt.Println("SuperMan.Sleep()...")
}

func (s *SuperMan) Sleep() {
	fmt.Println("SuperMan.Sleep()...")
}

// Fly 定义子类的新方法
func (s *SuperMan) Fly() {
	fmt.Println("SuperMan.Fly()...")
}

func (s *SuperMan) Print() {
	fmt.Println(s)
}

func main() {
	h := Human{"syk", "男"}
	h.Eat()
	h.Sleep()
	fmt.Println()

	// 定义一个子类对象
	// s := SuperMan{Human{"fyy", "女"}, 100} // 等价下面
	var s SuperMan
	s.Name = "hpq"
	s.Sex = "男"
	s.level = 90
	s.Eat()
	s.Sleep()
	s.Fly()
	s.Print()
}
