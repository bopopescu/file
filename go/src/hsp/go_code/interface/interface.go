package main

import (
	"fmt"
)

// 声明一个接口
type Usb interface {
	// 声明2个没有实现的方法
	connect()
	transport()
}

type Phone struct {
	Name string
}

// 让Phone实现usb接口的方法
func (p *Phone) connect() {
	fmt.Println("手机连接设备")
}
func (p *Phone) transport() {
	fmt.Println("手机传输文件")
}

type Camera struct {
}

// 让Camera实现usb接口的方法
func (p Camera) connect() {
	fmt.Println("相机连接设备")
}
func (p Camera) transport() {
	fmt.Println("相机传输文件")
}

// 计算机
type Computer struct {
}

func (c Computer) working(usb Usb) {
	usb.connect()
	usb.transport()
}

// 空接口
type empty interface {
}

func main() {
	// 不需要显示的指明实现的是哪个接口.  而且接口不能有变量!必须实现接口的所有方法.
	// 接口可以继承接口:内嵌匿名接口就可以了
	// 可以把任何变量赋给空接口,所有类型都实现了空接口.
	c := Computer{}
	var phone Usb = &Phone{} // 因为它是用指针类型实现这个接口的,所以要加&
	camera := Camera{}
	c.working(phone)
	c.working(camera)

	// 接口编程的最佳实践:sort排序,实现系统sort包内嵌的那3个方法:Len()...
}
