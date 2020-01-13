package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// 单精度 float32, 双精度 float64
	var f1 float32
	f1 = 12.1
	// 查看某个变量占用的字节大小
	fmt.Printf("f1的数据类型是%T	 f1占用的字节数是 %d   \n", f1, unsafe.Sizeof(f1))

	// 小数的一种写法
	f2 := .512
	fmt.Println("f2是", f2)

	// 科学计数法
	f3 := 5.1234e2
	f4 := 5.1234e-2
	fmt.Println("f3是", f3, "f4是", f4)

}
