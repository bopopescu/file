package utils

// 全局变量,要想被外面的文件使用同样要首字母大写
var Money int = 520

// 要想被外面的文件调用,方法名的首个字母必须大写

// 怎么把它打包成 .a静态库文件以便给同事使用
// 什么是编译器的逃逸分析
func Calculate(a float32,b float32) float32 {
	c := a +b
	return c
}   