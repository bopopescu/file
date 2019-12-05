package main

import (
	"fmt"
)
//感觉我们通过代码达成了修改字符串的过程，但真实的情况是：Go 语言中的字符串和其他高级语言（Java、C#）一样，默认是不可变的（immutable）。
//
//字符串不可变有很多好处，如天生线程安全，大家使用的都是只读对象，无须加锁；再者，方便内存共享，而不必使用写时复制（Copy On Write）等技术；字符串 hash 值也只需要制作一份。
//
//所以说，代码中实际修改的是 []byte，[]byte 在 Go 语言中是可变的，本身就是一个切片。
//
//在完成了对 []byte 操作后，在第 9 行，使用 string() 将 []byte 转为字符串时，重新创造了一个新的字符串。
//
//总结
//Go 语言的字符串是不可变的。
//修改字符串时，可以将字符串转换为 []byte 进行修改。
//[]byte 和 string 可以通过强制类型转换互转。

func main(){
	// string底层是一个byte数组,所以可以用切片来提取
	var str string = "woloveyou@mama"
	slice := str[10:]
	fmt.Println("slice=",slice)


	//如何修改字符串? 先把字符串转成byte切片修改,然后再转成string
	str1 := []byte(str)
	fmt.Printf("%p",str1)

	fmt.Println(str1,len(str1),len(str))
	str1[0]=121


	fmt.Println(str1)
	fmt.Printf("%p\n",str1)
	fmt.Printf("%p\n",&str1[0])
	fmt.Printf("%p\n",&str1[1])
	fmt.Printf("%p\n",&str1[5])

	fmt.Println(&str1)

	fmt.Println(string(str1))

	//str1[0]='z'
	fmt.Println("str1=",string(str1),str1)

	//如果有中文,
	str2 := []rune(str)
	str2[0]=38464211

	fmt.Println("str2=",str2)

	fmt.Printf("%p\n",&str2[0])
	fmt.Printf("%p\n",&str2[1])
	fmt.Printf("%p\n",&str2[2])
	fmt.Println(string(str2))

	return

	fmt.Println("str2=",str2,str1)

	fmt.Println("dasdas",str[2:6],str[2])
	str2[0]='北'
	fmt.Println("str=",str2[1])

	fmt.Printf("%c",str2[1])
	fmt.Println()

	str = string(str2)
	fmt.Println("str=",str)


}