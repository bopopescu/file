package main
import (
	"fmt"
)

func main() {

	//演示切片的使用 make
	var slice []float64 = make([]float64, 5, 100)
	fmt.Println(slice)

	var a bool = true
	fmt.Println(a)
	return
	//slice = make([]float64,10)
	slice[1] = 10
	slice[3] = 20
	//对于切片，必须make使用.
	fmt.Println(slice)
	fmt.Println("slice的size=", len(slice))
	fmt.Println("slice的cap=", cap(slice))
	fmt.Printf("slice的tpye=%T", slice)


	//方式3
	fmt.Println()
	//第3种方式：定义一个切片，直接就指定具体数组，使用原理类似make的方式
	var strSlice [3]string = [3]string{"tom", "jack", "mary"}
	//strSlice[3]="askyuy"
	fmt.Println("strSlice=", strSlice)
	fmt.Printf("slice的tpye=%T\n", strSlice)

	fmt.Println("strSlice size=", len(strSlice)) //3
	fmt.Println("strSlice cap=", cap(strSlice)) // ?

}