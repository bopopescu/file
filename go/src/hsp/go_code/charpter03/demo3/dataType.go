package main

import (
	"fmt"
	"unsafe"
)

func main() {

	// 基本数据类型:数值型(整数型和浮点型),字符型(byte存单个字母),布尔型,字符串

	// 有符号的类型
	var a int8
	a = -128
	fmt.Println("int8最小值是", a)

	// 无符号的类型
	var b uint8
	b = 0
	fmt.Println("uint8最小值是", b)

	// rune 类型 int32,用来存Unicode中文.  int就是int64

	// 查看某个数据的类型
	fmt.Printf("b的数据类型是%T   \n", b)
	// 查看某个变量占用的字节大小
	fmt.Printf("b的数据类型是%T	 b占用的字节数是 %d", b, unsafe.Sizeof(b))

	// 派生数据类型:指针,数组,结构体,管道,函数,切片,接口,map

}
