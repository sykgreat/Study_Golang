package HashMap

import (
	"testing"
)

var hashTable = NewHashMap[string, any]()

//func init() {
//	hashTable.Add("name", "张三")
//	hashTable.Add("age", 18)
//	hashTable.Add("sex", true)
//
//}

type User struct {
	Name   string
	Age    int
	Sex    bool
	Friend []User
}

func TestHashMap_Add(t *testing.T) {
	hashTable.Add("Information", User{
		Name:   "张三",
		Age:    18,
		Friend: nil,
	})
	t.Log(hashTable.Get("Information"))
}

func TestHashMap_Get(t *testing.T) {
	v := hashTable.Get("name")
	t.Log(v)
}

func TestHashMap_Remover(t *testing.T) {
	hashTable.Remove("name")
	v := hashTable.Get("name")
	t.Log(v)
}
