package main

import(
	"fmt"
)

func iterator1() {
	// 按字节来遍历的
	var str string = "hello,world"
	for i:=0;i<len(str);i++ {
		fmt.Printf("%c",str[i])
	}
}
func iterator2() {
	// for-range 遍历, 按字符来遍历的
	var str string = "nihao,world中国"
	for index,val := range str {
		fmt.Printf("index = %d,val = %c \n",index,val)
	}
}
func main(){
	// iterator1()
	iterator2()
	
	
	
	

}