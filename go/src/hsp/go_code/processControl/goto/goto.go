package main

import(
	"fmt"
)

func gotoDemo(){
	
	for i := 10;i>7;i-- {

		for j :=0;;j++{
			if j == 5{
				goto label
			}
			fmt.Println("数字",i)
		}
		
	}

	label:fmt.Println("程序结束啦")
}
func main(){
	gotoDemo()
	


}