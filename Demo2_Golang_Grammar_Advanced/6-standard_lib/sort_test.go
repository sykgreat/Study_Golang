package __standard_lib

import (
	"sort"
	"testing"
)

type user struct {
	ID   int
	Name string
	Age  int
}

type ById []user

func (b ById) Len() int {
	return len(b)
}

func (b ById) Less(i, j int) bool {
	return b[i].ID < b[j].ID
}

func (b ById) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func Test_Sort(t *testing.T) {

	list := []user{
		{ID: 4, Name: "nick4", Age: 21},
		{ID: 2, Name: "nick2", Age: 19},
		{ID: 1, Name: "nick", Age: 18},
		{ID: 3, Name: "nick3", Age: 20},
		{ID: 7, Name: "nick3", Age: 20},
		{ID: 5, Name: "nick3", Age: 20},
		{ID: 6, Name: "nick3", Age: 20},
	}

	sort.Slice(list, func(i, j int) bool {
		return list[i].Age < list[j].Age
	})
	t.Log(list)

	sort.Sort(ById(list))
	t.Log(list)
}
