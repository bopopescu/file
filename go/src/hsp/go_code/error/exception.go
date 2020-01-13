package main

import (
	"errors"
	"fmt"
)

func test() {
	defer func() {
		error := recover() //recover()内置函数用来捕获异常
		if error != nil {
			fmt.Println("---", error)
		}
	}()
	num1 := 100
	num2 := 0
	res := num1 / num2
	fmt.Println("", res)
}
func customerErr(name string) error {
	if name == "config.ini" {
		return nil
	} else {
		return errors.New("文件名错误")
	}
}
func main() {
	// 错误异常处理机制:使用defer+recover来捕获处理异常
	// test()
	// fmt.Println("错误处理机制")

	err := customerErr("abc.abc")
	if err != nil {
		panic(err)
	}
	fmt.Println("代码继续执行")

}
