package pointer

import (
	"fmt"
	// "strconv"
)

func main() {

	var i int = 10
	fmt.Println("i的地址是", &i)

	// ptr的类型是 *int
	// var ptr *int = &i

	// 基本数据类型:int,float,bool,string,数组和结构体都是值类型! 值类型作为参数传递时有值传递和址传递
	// 注意:在很多编程语言里面,数组是引用类型,但是go里面数组是值类型!所以就有值传统和址传统
	// 值类型都有对应的指针类型. 引用类型:指针,切片,map,管道,接口

	// 命名规范,标识符:
	var num_abc int
	fmt.Println("", num_abc)

}
