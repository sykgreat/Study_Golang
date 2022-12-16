package main

import "fmt"

type Driver interface {
	Drive(car Car)
}

type ZhangSan struct{}

func (zs ZhangSan) Drive(car Car) {
	fmt.Println("ZhangSan is driving...")
	car.Run()
}

type LiSi struct{}

func (ls LiSi) Drive(car Car) {
	fmt.Println("LiSi is driving...")
	car.Run()
}
