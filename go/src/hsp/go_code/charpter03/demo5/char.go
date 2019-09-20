package main

import (
	"fmt"
)

func main(){
	
	// 字符类型
	var c1 byte
	c1 = 'a'
	fmt.Println("c1是",c1)
	// 格式化输出
	fmt.Printf("c1是 %c  \n",c1)

	// 利用报错测试出汉字的Unicode编码. UTF-8编码是unicode编码的一种实现.
	c2 := '向'
	var c3 byte = '向'
	fmt.Printf("c2是 %c   \n",c2)
	fmt.Printf("c3是 %c",c3)



	
	
}