package main

import (
	"fmt"
	"reflect"
)

type resume struct {
	Name string `info:"Name" json:"name"`
	Sex  string `info:"Sex" json:"sex"`
}

func findTag(str interface{}) {
	t := reflect.TypeOf(str).Elem()

	for i := 0; i < t.NumField(); i++ {
		tagInfoString := t.Field(i).Tag.Get("info")
		tagJsonString := t.Field(i).Tag.Get("json")
		fmt.Println("info:", tagInfoString, "json: ", tagJsonString)
	}
}

func main() {
	var re resume
	findTag(&re)
}
