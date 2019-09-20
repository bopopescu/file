package main

import (
	"fmt"
)
// 给函数取别名
type myFunc func(int,int)(int)


func CalculatesumAndSub(a float32,b float32) (float32,float32) {
	c := a +b
	d := a-b
	return c,d
}   
// 定义返回值命名,直接写return就行了
func funcAsParameter(funcs func(int,int)(int),p1 int,p2 int) (res int) {
	res = funcs(p1,p2)
	return
}
// 这种写法你是不是很懵逼?
func funcAlias(funcs myFunc,p1 int,p2 int) (int) {
	res := funcs(p1,p2)
	return res
}
func test(a int,b int)(int){
	return a*b
}

// 可变参数
func sum(a int,args... int) int {
	sum := a;
	for i:=0;i<len(args);i++{
		sum += args[i]
	}
	return sum;
}
func main(){
	// 函数可以传值,也可以传地址
	c,d := CalculatesumAndSub(15,10)
	fmt.Println("运算和是",c," \t运算差是",d)
	
	// 只需要一个返回值
	e,_ := CalculatesumAndSub(100,10)
	fmt.Printf("运算和是%.2f \n",e)


	// 函数也是一种数据类型
	f := test
	fmt.Printf("f的类型是%T",f)
	res := f(12,12)
	fmt.Println("res是",res)

	//传递函数
	res1 := funcAsParameter(test,100,200)
	fmt.Println("res1是",res1)

	// 给int类型取别名, c语言的特性. 虽然都是int类型,但go认为他们是不同类型
	type Myint int
	var num Myint = 10
	fmt.Println("num是",num)

	res2 := funcAsParameter(test,100,200)
	fmt.Println("懵逼写法之 res2是",res2)

	res3 := sum(1,2,3)
	fmt.Println("可变参数之 res3是",res3)
}   