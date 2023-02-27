package TreeHashMap

import (
	"strconv"
	"testing"
)

var hashTable = NewTreeHashMap[string, any]()

//func init() {
//	hashTable.Add("a", 1)
//	hashTable.Add("b", 2)
//	hashTable.Add("c", 3)
//	hashTable.Add("d", 4)
//}

func TestTreeHashMap_Add(t *testing.T) {
	hashTable.Add("e", 5)
}

func TestTreeHashMap_Get(t *testing.T) {
	t.Log(hashTable.Get("a"))
	hashTable.Add("a", 6)
	t.Log(hashTable.Get("a"))
}

func TestTreeHashMap_Remover(t *testing.T) {
	t.Log(hashTable.Get("a"))
	t.Log(hashTable.Remove("a"))
	t.Log(hashTable.Get("a"))
}

func TestTreeHashMap_Performance_Insert(t *testing.T) {
	for i := 0; i < 1000000; i++ {
		str := "value:" + string(rune(i))
		hashTable.Add(string(rune(i)), str)
	}
}

func init() {
	for i := 0; i < 1000000; i++ {
		str := "value:" + strconv.Itoa(i)
		hashTable.Add(strconv.Itoa(i), str)
	}
}

func TestTreeHashMap_Performance_Search(t *testing.T) {
	t.Log(hashTable.Get("10000"))
}
