package main

import "fmt"

func main()  {
	var intArr [3]int //int占8个字节
	//当我们定义完数组后，其实数组的各个元素有默认值 0
	fmt.Println(intArr)
	intArr[0] = 10
	intArr[1] = 20
	intArr[2] = 30
	fmt.Println(intArr)
	fmt.Printf("intArr的地址=%p intArr[0] 地址%p intArr[1] 地址%p intArr[2] 地址%p\n",
		&intArr, &intArr[0], &intArr[1], &intArr[2])


	intArr[1]=33
	fmt.Println(intArr)


}