package main

import "fmt"

type AbstractApple interface {
	ShowApple()
}

type AbstractBanana interface {
	ShowBanana()
}

type AbstractPineapple interface {
	ShowPineapple()
}

type AbstractFactory interface {
	CreatApple() AbstractApple
	CreatBanana() AbstractBanana
	CreatPineapple() AbstractPineapple
}

type ChinaApple struct{}

func (ca *ChinaApple) ShowApple() {
	fmt.Println("i am china apple")
}

type ChinaBanana struct{}

func (cb *ChinaBanana) ShowBanana() {
	fmt.Println("i am china banana")
}

type ChinaPineapple struct{}

func (cp *ChinaPineapple) ShowPineapple() {
	fmt.Println("i am china pineapple")
}

type ChinaFactory struct{}

func (cf *ChinaFactory) CreatApple() AbstractApple {
	return new(ChinaApple)
}

func (cf *ChinaFactory) CreatBanana() AbstractBanana {
	return new(ChinaBanana)
}

func (cf *ChinaFactory) CreatPineapple() AbstractPineapple {
	return new(ChinaPineapple)
}

type USAApple struct{}

func (ca *USAApple) ShowApple() {
	fmt.Println("i am usa apple")
}

type USABanana struct{}

func (cb *USABanana) ShowBanana() {
	fmt.Println("i am usa banana")
}

type USAPineapple struct{}

func (cp *USAPineapple) ShowPineapple() {
	fmt.Println("i am usa pineapple")
}

type USAFactory struct{}

func (uf *USAFactory) CreatApple() AbstractApple {
	return new(USAApple)
}

func (uf *USAFactory) CreatBanana() AbstractBanana {
	return new(USABanana)
}

func (uf *USAFactory) CreatPineapple() AbstractPineapple {
	return new(USAPineapple)
}

func main() {
	chinaFactory := new(ChinaFactory)

	chinaApple := chinaFactory.CreatApple()
	chinaApple.ShowApple()

	chinaBanana := chinaFactory.CreatBanana()
	chinaBanana.ShowBanana()

	chinaPineapple := chinaFactory.CreatPineapple()
	chinaPineapple.ShowPineapple()

	usaFactory := new(USAFactory)

	usaApple := usaFactory.CreatApple()
	usaApple.ShowApple()

	usaBanana := usaFactory.CreatBanana()
	usaBanana.ShowBanana()

	usaPineapple := usaFactory.CreatPineapple()
	usaPineapple.ShowPineapple()
}
