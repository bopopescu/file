package main

import (
	"fmt"
)

func breakDemo1() {
	for i := 10; ; i-- {
		if i == 5 {
			break
		}
		fmt.Println("数字", i)
	}
}
func breakLabel() {
label:
	for i := 10; i > 7; i-- {

		for j := 0; ; j++ {
			if j == 5 {
				// break
				break label

			}
			fmt.Println("数字", i)
		}

	}
}
func main() {
	// break是跳出离它最近的那一层循环!
	// breakDemo1()
	// break跳到某个标签
	breakLabel()

}
