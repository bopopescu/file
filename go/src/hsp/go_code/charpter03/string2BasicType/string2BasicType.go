package main

import (
	"fmt"
	"strconv"
)

func main(){


	// string转成其他类型
	var str string = "true"
	// 这个函数返回2个值,但是另外一个值我不关心,所有用_忽略
	var b bool
	b,_ = strconv.ParseBool(str)
	fmt.Println("转换结果",b)


	var str2 string = "1223"
	var r int64 
	r,_ = strconv.ParseInt(str2,10,64)
	fmt.Printf("类型  %T, 值 %v",r,r)
    
}