package main

import (
	"fmt"
	"strconv"
)

func main(){
	
	// 必须显示的转换
	var a int = 32
	var b float32 = float32(a)
	fmt.Printf("b=%f",b)

	var c int = 32
	// var d float32 = 30.0
	var e bool = false
	// var f byte = 'h'

	// 调用api将基本类型转换成字符串
	var str string = fmt.Sprintf("%d",c)
	var str1 string = fmt.Sprintf("%t",e)
	
	fmt.Println("",str)
	fmt.Println("",str1)
	
	
	// 调用api将基本类型(int系列,float系列,bool,string,数组和结构体struct)转换成字符串
	var str11 string = strconv.FormatInt(int64(c),10)
	var str12 string = strconv.FormatFloat(float64(c),'f',10,64)
	
	fmt.Println("",str11)
	fmt.Println("",str12)
	
	// strconv.itoa

	
	
}