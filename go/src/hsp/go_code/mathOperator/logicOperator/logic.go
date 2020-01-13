package main

import (
	"fmt"
)

func test() bool {
	fmt.Println("测试函数被调用了")
	return true
}
func main() {
	var a, b int
	a = 2
	b = 4
	// 逻辑 与(也叫短路与)  或(短路或)   非
	if a > b && a > 0 {
		fmt.Println(a > b)
	}

	if a > 10 || test() {
		fmt.Println("结果为真")
	}

}
