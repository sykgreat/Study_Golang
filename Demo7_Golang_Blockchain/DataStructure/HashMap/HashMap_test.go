package HashMap

import "testing"

var hashTable = CreateHashMap[string, any]()

func init() {
	hashTable.AddKeyValue("name", "张三")
	hashTable.AddKeyValue("age", 18)
	hashTable.AddKeyValue("sex", true)
}

func TestHashMap_AddKeyValue(t *testing.T) {
	hashTable.AddKeyValue("Information", User{
		Name: "张三",
		Age:  18,
		Sex:  true,
		Friend: []User{
			{
				Name: "李四",
				Age:  19,
				Sex:  true,
			},
			{
				Name: "王五",
				Age:  20,
				Sex:  false,
			},
		},
	})
}

func TestHashMap_GetValueForKey(t *testing.T) {
	hashTable.AddKeyValue("name", "张三")
	v := hashTable.GetValueForKey("name")
	t.Log(v)
}

func TestHashMap_RemoverValueForKey(t *testing.T) {
	hashTable.RemoverValueForKey("name")
	v := hashTable.GetValueForKey("name")
	t.Log(v)
}
