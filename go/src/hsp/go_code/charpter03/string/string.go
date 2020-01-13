package main

import (
	"fmt"
)

func main() {

	// 字符串类型
	var str1 string
	str1 = "中国"
	fmt.Println("字符串", str1)

	// 字符串一旦赋值了,就不能再去修改它的值

	// 字符串的2种写法:双方号和反引号:键盘最左边那个符号,可以原生形式输出.
	str2 := `  原生形式输出文本"这个功能好酷啊!"  `
	fmt.Println("", str2)

	str3 := "拼接字符串" + str2
	fmt.Println("", str3)

	// 多个字符串拼接时要注意,加号要放在行位,要不然golang会自动给你加上分号,然后报错.
	str4 := "拼接字符串" + str2 +
		"拼接字符串" + str2 + "拼接字符串" +
		str2
	fmt.Println("", str4)

}
