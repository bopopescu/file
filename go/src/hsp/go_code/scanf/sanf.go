package main

import (
	"fmt"
)

func main() {
	// 方式1
	var name string
	var age int
	var salary float32
	var isPass bool

	fmt.Println("请输入姓名:")
	// 等待用户输入,必须回车才能结束输入
	fmt.Scanln(&name)

	fmt.Println("请输入年龄:")
	fmt.Scanln(&age)

	fmt.Println("请输入薪水:")
	fmt.Scanln(&salary)

	fmt.Printf("名字是 %v \n 年龄是 %v \n 薪水是 %v \n ", name, age, salary)

	// 第二种方式:一次性输入
	fmt.Scanf("%s %d %f %t", &name, &age, &salary, &isPass)
}
