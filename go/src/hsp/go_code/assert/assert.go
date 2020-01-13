package main

import (
	"fmt"
)

type Point struct {
	x int
	y int
}

func main() {
	// 空接口
	var a interface{}
	var point Point = Point{1, 2}
	a = point

	var b Point
	// b = a  //报错
	// 必须用类型断言,才不会报错. 因为编译器不知道变量a是什么类型.
	b = a.(Point)
	fmt.Println("", b)

	// 类型断言用来判断 a是否可以转成Point类型.
	b, ok := a.(Point)
	if ok == true {
		fmt.Println("类型断言转换成功", b)
	}

	// 类型断言的结果可以结合switch语句来使用
}
