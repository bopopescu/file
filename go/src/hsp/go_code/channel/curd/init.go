package main

import(
	"fmt"
)
type Cat struct{
	Name string
	Age int
}

func channelCreate1(){
	// 管道声明为只写  var intchan chan<- int
	// 管道声明为只读  var intchan <-chan int
	var intchan chan int

	intchan = make(chan int,3)

	fmt.Printf("intchan管道的值是%v,intchan管道的地址是%p \n",intchan,&intchan)

	// 向管道写入同种数据类型, FIFO 先进先出
	intchan <-10
	num:=123
	intchan <-num
	fmt.Printf("intchan管道的长度是%v,intchan管道的容量是%v \n",len(intchan),cap(intchan))

	// 从管道读数据,长度会减少
	 num2 := <- intchan
	fmt.Printf("读取的数据是%v,intchan管道的长度是%v,intchan管道的容量是%v \n",num2,len(intchan),cap(intchan))
}
func channelCreate2(){
	// 如果想要向管道写入各种数据类型,必须用空接口
	var intchan chan interface{}
	intchan = make(chan interface{},3)

	fmt.Printf("intchan管道的值是%v,intchan管道的地址是%p \n",intchan,&intchan)

	// 向管道写入各种数据类型, FIFO 先进先出
	intchan <-"xss"
	num:=123
	intchan <-num
	intchan <-Cat{"kitty",3}
	fmt.Printf("intchan管道的长度是%v,intchan管道的容量是%v \n",len(intchan),cap(intchan))

	// 从管道读数据读出来的是接口类型,因为类型是接口,所有容易出问题,需要用类型断言
	<- intchan
	<- intchan
	cat2ha := <- intchan
	a := cat2ha.(Cat)
	fmt.Printf("读取的数据是%v,intchan管道的长度是%v,intchan管道的容量是%v \n",cat2ha,len(intchan),cap(intchan))
	fmt.Printf("小猫的名字是%v",a.Name)
}
func main(){
	// channelCreate1()
	channelCreate2()

	
}