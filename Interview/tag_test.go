package Interview

import (
	"fmt"
	"reflect"
	"testing"
)

type TagTest struct {
	Name      string `bilibili:"bilibili_name" json:"name"`
	PublicMax string `bilibili:"bilibili_publicMax" json:"publicMax"`
}

func PrintTag(ptr any) {
	of := reflect.TypeOf(ptr)
	if of.Kind() != reflect.Ptr || of.Elem().Kind() != reflect.Struct {
		panic("传入的参数不是结构体指针")
		return
	}
	v := reflect.ValueOf(ptr).Elem()
	fmt.Println(v)
	for i := 0; i < v.NumField(); i++ {
		tag := v.Type().Field(i).Tag
		giligili := tag.Get("bilibili")
		fmt.Println(giligili)
	}
}

func Test_Tag(t *testing.T) {
	tag := TagTest{
		Name:      "bilibili",
		PublicMax: "100",
	}
	PrintTag(&tag)
}
