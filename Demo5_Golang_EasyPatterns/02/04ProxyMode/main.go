package main

import "fmt"

type Goods struct {
	Name         string
	Authenticity bool
}

type Shopping interface {
	Buy(goods *Goods)
}

type USAShopping struct{}

func (u *USAShopping) Buy(goods *Goods) {
	fmt.Println("在美国代购了，", goods.Name)
}

type JapanShopping struct{}

func (j *JapanShopping) Buy(goods *Goods) {
	fmt.Println("在日本代购了，", goods.Name)
}

type OverseasProxy struct {
	shopping Shopping
}

func (op *OverseasProxy) Buy(goods *Goods) {
	if op.Distinguish(goods) {
		op.shopping.Buy(goods)
		op.Check(goods)
	}
}

func NewProxy(shopping Shopping) Shopping {
	return &OverseasProxy{
		shopping: shopping,
	}
}

func (op *OverseasProxy) Distinguish(goods *Goods) bool {
	fmt.Println("鉴定", goods.Name, "商品真伪")
	if goods.Authenticity == false {
		fmt.Println("假货")
		return false
	}
	fmt.Println("真货")
	return goods.Authenticity
}

func (op *OverseasProxy) Check(goods *Goods) {
	fmt.Println("对", goods.Name, "进行海关检查")
}

func main() {
	g1 := &Goods{
		Name:         "手办",
		Authenticity: true,
	}

	g2 := &Goods{
		Name:         "CET4",
		Authenticity: false,
	}

	usa := new(USAShopping)
	japan := new(JapanShopping)

	usaProxy := NewProxy(usa)
	japanProxy := NewProxy(japan)

	japanProxy.Buy(g1)
	usaProxy.Buy(g2)
}
