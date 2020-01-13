package main

import (
	"fmt"
)

func test(n1 int, n2 int) int {
	// 创建资源文件时会经常使用这个,主要用来释放资源!例如创建数据库连接
	// 当执行到defer语句时,暂时不执行,先压入defer栈里面. 变量值也会入栈
	// 当函数执行完毕后,再弹出栈时,再执行
	defer fmt.Println("n1的值是", n1)
	defer fmt.Println("n2的值是", n2)

	n1++
	n2++
	return n1 + n2

}

func main() {

	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	n1 := 10
	n2 := 0
	res := n1 / n2
	//res := test(5,10);
	fmt.Println("res的值是", res)

}
