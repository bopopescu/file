package main

import (
	"fmt"
)

func continueLabel() {
label:
	for i := 10; i > 7; i-- {

		for j := 0; ; j++ {
			if j == 5 {
				// continue
				continue label
			}
			fmt.Println("数字", i)
		}

	}
}
func main() {
	// continue是跳出离它最近的那一层循环! 最近的那个循环体里面在continue后面的内容不会执行了
	continueLabel()

}
