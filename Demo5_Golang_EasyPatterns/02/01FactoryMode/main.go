package main

import "fmt"

type Fruit interface {
	Show()
}

type AbstractFactory interface {
	CreatFruit() Fruit
}

type Apple struct{}

func (apple *Apple) Show() {
	fmt.Println("i am apple")
}

type Banana struct{}

func (banana *Banana) Show() {
	fmt.Println("i am banana")
}

type Pineapple struct{}

func (pineapple *Pineapple) Show() {
	fmt.Println("i am pineapple")
}

type AppleFactory struct{}

func (appleFactory *AppleFactory) CreatFruit() Fruit {
	return new(Apple)
}

type BananaFactory struct{}

func (bananaFactory *BananaFactory) CreatFruit() Fruit {
	return new(Banana)
}

type PineappleFactory struct{}

func (pineappleFactory *PineappleFactory) CreatFruit() Fruit {
	return new(Pineapple)
}

func main() {
	appleFactory := new(AppleFactory)
	apple := appleFactory.CreatFruit()
	apple.Show()

	bananaFactory := new(BananaFactory)
	banana := bananaFactory.CreatFruit()
	banana.Show()

	pineappleFactory := new(PineappleFactory)
	pineapple := pineappleFactory.CreatFruit()
	pineapple.Show()
}
