package main

import "fmt"

type Hero struct {
	Name string
	Wife string
}

func (hero *Hero) Shwo() {
	fmt.Printf("Hero is %v, hero's wife is %v\n", hero.Name, hero.Wife)
}

func (hero *Hero) GetName() string {
	return hero.Name
}

func (hero *Hero) SetName(newName string) {
	hero.Name = newName
}

func (hero *Hero) GetWife() string {
	return hero.Wife
}

func (hero *Hero) SetWife(newWife string) {
	hero.Wife = newWife
}

func main() {
	hero := Hero{Name: "syk", Wife: "fyy"}
	hero.Shwo()
}
