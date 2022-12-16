package main

import "fmt"

// Animal 本质是指针
type Animal interface {
	Sleep()
	GetColor() string // 获取动物的颜色
	GetType() string  // 获取动物的类型
}

// Cat 定义一个具体的类
type Cat struct {
	Color string // 猫的颜色
}

func (cat *Cat) Sleep() {
	fmt.Println("Cat is sleeping...")
}

func (cat *Cat) GetColor() string {
	return cat.Color
}

func (cat *Cat) GetType() string {
	return "Cat"
}

// Dog 定义一个具体的类
type Dog struct {
	Color string // 狗的颜色
}

func (cat *Dog) Sleep() {
	fmt.Println("Dog is sleeping...")
}

func (cat *Dog) GetColor() string {
	return cat.Color
}

func (cat *Dog) GetType() string {
	return "Dog"
}

func ShowAnimal(animal Animal) {
	animal.Sleep()
	fmt.Println("color = ", animal.GetColor())
	fmt.Println("type = ", animal.GetType())
}

func main() {
	var cat Animal // 接口的数据类型，父类指针
	cat = &Cat{"Green"}
	cat.Sleep()
	ShowAnimal(cat)
	fmt.Println()

	var dog Animal
	dog = &Dog{"Yellow"}
	dog.Sleep()
	ShowAnimal(dog)
	fmt.Println()
}
