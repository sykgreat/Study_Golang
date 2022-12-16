package test

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"testing"
)

var user *User

func Test1(t *testing.T) {
	idx := rand.Intn(10)
	user = &User{
		Name: "zhangsan",
		Age:  18,
		Sex:  1,
		class: &Class{
			CName: "class-" + strconv.Itoa(idx),
			Index: uint(idx),
		},
	}
	fmt.Println(user)
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, os.Kill)
	s := <-c
	fmt.Println("receive signal -> ", s)
}

type User struct {
	Name  string
	Age   uint8
	Sex   uint8
	class *Class
}

type Class struct {
	CName string
	Index uint
}
