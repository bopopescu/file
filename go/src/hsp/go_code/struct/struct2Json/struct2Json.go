package main

import (
	"encoding/json"
	"fmt"
)

type Point struct {
	x int
	y int
}
type Rect struct {
	leftUp, rightDown Point
}

type Ghost struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Sex  string `json:"sex"`
}

func demo1() {
	r1 := Rect{Point{1, 1}, Point{4, 4}}
	fmt.Println("", r1)
}
func demo2() {
	ghost := Ghost{"zobie", 100, "male"}
	jsonstr, err := json.Marshal(ghost)
	if err != nil {
		fmt.Println("错误", err)
	}
	fmt.Println("解析后结果是", string(jsonstr))
}
func main() {
	// 相同结构的的结构体之间可以强转
	// demo1()
	demo2()
}
