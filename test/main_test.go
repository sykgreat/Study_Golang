package test

import (
	"reflect"
	"testing"
)

func Test_reflect(t *testing.T) {
	//t.Log(GetValueSize(true))
	t.Log(GetValueSize(map[string]interface{}{
		"1": 123,
		"2": "456",
		//"3": "789",
		//"4": []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0},
	}))
	//t.Log(GetValueSize([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}))
	//t.Log(GetValueSize([]string{"1", "2"}))
	//t.Log(GetValueSize(123))
	//t.Log(GetValueSize("123"))
	t.Log(GetValueSize(
		users{
			Name: "123",
			Age:  123,
			Sex:  true,
			Bzd: map[string]interface{}{
				"1": 123,
				"2": "456",
			},
		},
	))
}

func GetValueSize(value interface{}) int64 {
	var size int64 = 0
	vo := reflect.ValueOf(value)
	switch vo.Kind() {
	case reflect.Map:
		for _, key := range vo.MapKeys() {
			size += GetValueSize(key.Interface())
			size += GetValueSize(vo.MapIndex(key).Interface())
		}
	case reflect.Slice:
		for i := 0; i < vo.Len(); i++ {
			size += GetValueSize(vo.Index(i).Interface())
		}
	case reflect.Struct:
		for i := 0; i < vo.NumField(); i++ {
			size += GetValueSize(vo.Field(i).Interface())
		}
	}
	size += int64(vo.Type().Size())
	return size
}

type users struct {
	Name string
	Sex  bool
	Age  int
	Bzd  map[string]interface{}
}
