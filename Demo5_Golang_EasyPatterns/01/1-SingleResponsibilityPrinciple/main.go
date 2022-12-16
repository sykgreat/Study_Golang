package main

import "fmt"

type ClothesShop struct{}

func (cs *ClothesShop) Style() {
	fmt.Println("买东西的装扮")
}

type ClothesWork struct{}

func (cw *ClothesWork) Style() {
	fmt.Println("上班的装扮")
}

func main() {
	cs := &ClothesShop{}
	cs.Style()

	cw := &ClothesWork{}
	cw.Style()
}
