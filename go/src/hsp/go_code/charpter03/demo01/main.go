package main

import "fmt"

// 全局变量
var globalVarity int = 100

// 全局变量也可以写成下面的形式
var {
	globalVarity1 = 200
}

func main(){
	// 声明
	var i int
	i = 10
	fmt.Println("i=",i)


	// 类型推导
	var num = 10.1
	fmt.Println("i=",num);


	// 简写, 不能在函数体外或者全局变量里面写
	abc := "i am a men"
	fmt.Println("abc =",abc)


	// 一次性声明多个变量
	var n1,n2,n3,n4 int
	fmt.Println("n1值是:",n1,"n2值是:",n2,"n3值是:",n3,"n4值是:",n4)
	
	输出全局变量
	fmt.Println("全局变量globalVarity是",globalVarity)
	fmt.Println("全局变量globalVarity1是",globalVarity1)
	
}