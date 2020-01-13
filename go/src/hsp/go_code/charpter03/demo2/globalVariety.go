package main

import "fmt"

// 全局变量
var globalVarity int = 100

// 全局变量也可以写成下面的形式
var (
	globalVarity1 = 200
)

func main() {

	// 输出全局变量
	fmt.Println("全局变量globalVarity是", globalVarity)
	fmt.Println("全局变量globalVarity1是", globalVarity1)

}
