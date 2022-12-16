package main

import "fmt"

type car struct{}

func (c *car) sleep() {
	fmt.Println("car sleeps...")
}

type blackCar struct {
	//c *car
}

func (bc *blackCar) sleep(c *car) {
	c.sleep()
}

func (bc *blackCar) eat() {
	fmt.Println("black car eats...")
}

func main() {
	c := &car{}
	c.sleep()
	fmt.Println("----------------------------------------")
	bc := &blackCar{}
	//bc.c.sleep()
	bc.sleep(c)
	bc.eat()
}
