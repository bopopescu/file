package main

import (
	"fmt"
)

type Person struct {
	Name string
}

// 给Person绑定一个方法
func (p Person) doing() {
	fmt.Println("名字是", p.Name)
}

// 给Person绑定一个方法
func (p Person) calculate(a int, b int) {
	res := a + b
	fmt.Println("和是", res)
}

// 给*Person绑定一个方法
func (p *Person) calculate11(a int, b int) {
	res := a + b
	fmt.Println("地址传递", res)
}

// 给自定义的类型绑定方法
type integer int

func (i integer) print() {
	fmt.Println("给int绑定方法", i)
}

// 实现String()方法
func (i integer) String() string {
	return "相当于java的toString()方法"
}

func main() {
	// 方法是作用在指定的数据类型上面的! 自定义类型(type)也可以有方法,而不仅仅是结构体才能有方法!
	//func (p Person)doing()表示Person结构体有一个方法doing.
	var p Person
	p.doing()
	p.calculate(5, 10)
	(&p).calculate(5, 10)   // 值拷贝
	(&p).calculate11(5, 10) // 地址传递

	// 调用fmt.Println(i)方法时默认会调用i的String()方法
	var i integer = 88
	fmt.Println(i)
}
