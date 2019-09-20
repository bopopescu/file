package main

import(
	"fmt"
)

func test(a int) int{
	return a-9;
}
func main(){

	if sum :=38; sum>40 {
		fmt.Println("你的年龄大于20")
	}else{

	}

	// 第一种switch形式
	var a int = 10
	switch test(a) {
	case 1:fmt.Println("a值为1")
	case 2:fmt.Println("a值为2")
	case 3:fmt.Println("a值为3")

	case 4,5,6,7,8,9:fmt.Println("a值为4")
	
	default:fmt.Println("a其他值")
	}

	// 第二种switch形式
	var b int = 10
	switch {
		case b <= 10:fmt.Println("b值不大于10")
		case b > 10:fmt.Println("b值大于10")
		case b > 20:fmt.Println("b值大于20")
		default:fmt.Println("b其他值")
	}

	// 第3种switch形式
	switch c:= 10; {
		case c <= 10:fmt.Println("c值不大于10")
		case c > 10:fmt.Println("c值大于10")
		case c > 20:fmt.Println("c值大于20")
		default:fmt.Println("c其他值")
	}

	// switch的穿透:fallthrough,默认穿透一层
	var d int = 10
	switch d{
		case 10:fmt.Println("d值10")
				fallthrough
		case 20:fmt.Println("d值20")
		default:fmt.Println("d其他值")
	}
}