package main

// import "fmt"
// 包路径默认是从gopath环境变量下面的src下面开始的
// main包只能有一个.
// import "go_code/package/utils"
import (
	"fmt"
	// 还可以给包取别名
	// app "go_code/package/utils"
	"go_code/package/utils"
)


func main(){
	i := utils.Calculate(5,10)
	fmt.Printf("运算结果是%.2f \n",i)

	fmt.Println("全局变量是",utils.Money)
	
}   