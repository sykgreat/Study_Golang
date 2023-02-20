package test

import (
	"fmt"
	"testing"
)

func Any[T any](expr bool, a, b T) T {
	if expr {
		return a
	}
	return b
}

type Person struct {
	gender int
}

func TestT(t *testing.T) {
	var p *Person
	genderDesc := Any(p == nil, "未知", Any(p.gender == 1, "男", "女")) // panic
	fmt.Println(genderDesc)
}

func TestName(t *testing.T) {
	Watch[B]()
}

func Watch[T interface{ Foo() }]() (x *T) {
	switch any(x).(type) {
	case A:
		{
			return any(&A{}).(*T)
		}
	case B:
		{
			return any(&B{}).(*T)
		}
	default:
		{
			panic("unknown type")
		}
	}
}

type A struct{}

func (a A) Foo() {}

type B struct{}

func (b B) Foo() {}
