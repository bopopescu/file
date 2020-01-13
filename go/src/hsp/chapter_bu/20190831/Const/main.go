package main

import "fmt"

func main() {

	var num int

	const name string = "123"

	//必须初始化

	//不可修改

	//only bool int string

	fmt.Println(name, num)

	const (
		a = 1
		b = 2
	)
	fmt.Println(a, b)

	const (
		d = iota
		e
		f
	)

	//依然访问范围 大写可以全局访问
	fmt.Println(d, e, f)
}
