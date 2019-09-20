package main

import (
	"fmt"
)

func main(){
	// string底层是一个byte数组,所以可以用切片来提取
	var str string = "woloveyou@mama"
	slice := str[10:]
	fmt.Println("slice=",slice)

	//如何修改字符串? 先把字符串转成byte切片修改,然后再转成string
	str1 := []byte(str)
	str1[0]='z'
	fmt.Println("str1=",str1)

	//如果有中文,
	str2 := []rune(str)
	str2[0]='北'
	str = string(str2)
	fmt.Println("str=",str)


}