package main

import "fmt"

type Car interface {
	Run()
}

type Benz struct{}

func (b *Benz) Run() {
	fmt.Println("Benz is running...")
}

type BMW struct{}

func (b *BMW) Run() {
	fmt.Println("BMW is running...")
}
