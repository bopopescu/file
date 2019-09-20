package main

import (
	"fmt"
	"strings"
)

func adder() func(int) int{
	// p只在第一次调用的时候初始化,后面再调用就不会初始化了.为什么?
	// 返回的函数和外面的变量p共同构成一个闭包
	var p int = 10
	return func (k int)int{
		p = p+k
		return p
	}
}

func makeSuffix(suffix string) func(string)string{
	return func(name string)string{
		if !strings.HasSuffix(name,suffix){
			return name+suffix
		}
		return name
	} 
}
func main(){
	// var haha = adder
	// var f = haha()

	f1 := adder()
	fmt.Println("闭包结果",f1(1))
	fmt.Println("闭包结果",f1(2))

	// 闭包的最佳实践
	// 闭包的好处:它可以保留上次引用的某个值,下一次继续使用,比如这里的jpg文件后缀.
	file := makeSuffix(".jpg")
	fmt.Println("文件处理后的结果",file("girl"))

}