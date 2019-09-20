package main

import (
	"fmt"
	
)

// 定义一个猫的结构体
type Cat struct{
	Name string
	Age int
	Color string
}
func structCreat1(){
	var cat1 Cat
	cat1.Name = "小白"
	cat1.Age = 1
	cat1.Color= "white"

	cat2 := cat1
	cat2.Name = "大白"
	fmt.Println("cat1是",cat1)
	fmt.Println("cat2是",cat2)
}

func structCreat2(){
	var cat1 Cat = Cat{"小狗",1,"gray"}
	var cat2 Cat = Cat{Age:1,Color:"gray",Name:"小狗"}
	fmt.Println("cat1是",cat1)
}
func structCreat3(){
	var cat1 *Cat = new(Cat)
	(*cat1).Name = "大哈"
	cat1.Name = "小哈"
	// go的设计者觉得这个写法麻烦,默认在底层做了简化
	fmt.Println("cat1是",cat1)
}
func structCreat4(){
	var cat1 *Cat = &Cat{"小狗",1,"gray"}
	(*cat1).Name = "大哈"
	cat1.Name = "小哈"
	// go的设计者觉得这个写法麻烦,默认在底层做了简化
	fmt.Println("cat1是",cat1)
}
func main(){
	// 结构体是值类型,而不是c语言里面的引用类型!值类型只是值拷贝,并不会影响原来的!!以下举例告诉你为什么
	// 结构体的属性成员在内存里面是连续分配地址的
	//structCreat1()

	// structCreat2()

	// 指针来指向结构体
	// structCreat3()

	structCreat4()
}