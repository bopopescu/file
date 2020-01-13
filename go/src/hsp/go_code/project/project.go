package main

import (
	"fmt"
)

// 家庭收支项目
func main() {

	// 接受用户输入的选项
	key := ""
	// 是否退出显示界面
	loop := true

	fmt.Println("-------------家庭收支记账软件------------")
	fmt.Println("		1 收支明细")
	fmt.Println("		2 登记收入")
	fmt.Println("		3 登记支出")
	fmt.Println("		4 退出软件")

	for {
		fmt.Scanln(&key)
		switch key {
		case "1":
			fmt.Println("-------------收支明细------------")
		case "2":
			fmt.Println("-------------登记收入------------")
		case "3":
			fmt.Println("-------------登记支出------------")
		case "4":
			loop = false
		}

		if loop != true {
			break
		}
	}
}
