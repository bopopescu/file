package main

import (
	"fmt"
)

var (
	// 全局匿名函数,大写.
	GlobalFun = func(a int, b int) (res int) {
		return a * b
	}
)

func main() {
	// 第一种写法:定义匿名函数的时候就调用了
	res := func(p1 int, p2 int) int {
		return p1 + p2
	}(15, 20)
	fmt.Println("匿名函数的结果是", res)

	// 第二种写法
	v := func(p1 int, p2 int) int {
		return p1 + p2
	}
	res1 := v(11, 11)
	fmt.Println("匿名函数的结果是", res1)

	res2 := GlobalFun(2, 5)
	fmt.Println("全局匿名函数的结果是", res2)

}
