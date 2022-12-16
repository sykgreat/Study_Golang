package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title  string   `info:"Title" json:"title"`
	Year   int      `info:"Year" json:"year"`
	Price  int      `info:"Price" json:"price"`
	Actors []string `info:"Actors" json:"actors"`
}

func main() {
	movie := Movie{"喜剧之王", 2000, 10, []string{"周星驰", "不知道"}}

	// 序列化 struct => json
	jsonStr, err := json.Marshal(movie)
	if err != nil {
		fmt.Println("json marshal error:", err)
		return
	}
	fmt.Printf("jsonStr = %s\n", jsonStr)

	// 反序列化 json => struct
	myMovie := Movie{}
	err = json.Unmarshal(jsonStr, &myMovie)
	if err != nil {
		fmt.Println("josn unmarshal error: ", err)
		return
	}
	fmt.Print("%v\n", myMovie)
}
