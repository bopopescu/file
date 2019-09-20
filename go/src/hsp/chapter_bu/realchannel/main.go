package main

import "fmt"

var mychan chan int
type Cat struct {
	Name string
	Age int
}
func assert(){

	var chani chan interface{}

	chani = make(chan interface{},3)

	var newcat = Cat{
		Name:"sdds",
		Age:13,
	}
	chani<-newcat
	chani<-12

	var testNewcat =<-chani

	fmt.Println(testNewcat)


	a:= testNewcat.(Cat)
	fmt.Println(a)

	fmt.Println(a.Name)
	fmt.Printf("%T,%T,%v,%v",testNewcat,a,testNewcat,a)


}

//特别注意close 关闭 channel信道   之后再行循环 for range  / for ok
func closetest(){
	intchan:= make(chan int ,2)
	intchan<-123
	intchan<-24

	//close(intchan)

	//for v:=range intchan{
	//	fmt.Println(v)
	//}
	//return

	i,j:=<-intchan
	fmt.Println(i,j)
	i2,j2:=<-intchan
	fmt.Println(i2,j2)
	close(intchan)
	i21,j21:=<-intchan
	fmt.Println(i21,j21)
return
	close(intchan)

	//intchan<-2423

	n1:=<-intchan
	fmt.Println("sda")
	fmt.Println(n1)


	//for range

	intchan2:=make(chan int,100)
	for i:=0;i<100;i++{
		intchan2<-i*2
	}

	close(intchan2)
	for v:=range intchan2{
		fmt.Println("v=",v)
	}


}




func write(goch1 chan int){

	for i:=1;i<50;i++{
		goch1<-i
		fmt.Println("write data is",i)

	}
	close(goch1)
}
func read(goch1 chan int,goch2 chan bool){

	for{
		v,ok:=<-goch1
		if !ok{
			break
		}
		fmt.Println("read data is",v)
	}

	goch2<-true
	close(goch2)
}


func main() {
	//closetest()
	//
	//return

	//管道 协程
	goch1:=make(chan int ,50)
	goch2:=make(chan bool ,1)

	////
	//goch3:=make(chan int ,6)
	//
	//for i:=0;i<6;i++{
	//	goch3<- i*3
	//}
	//close(goch3)
	//
	//for{
	//	v,ok:=<-goch3
	//	if !ok{
	//		fmt.Println("666666666")
	//
	//		break
	//	}
	//	fmt.Println("read data is",v)
	//}
	////close(goch3)
	////v,ok:=<-goch3
	////fmt.Println(v,ok)
	//return


	go write(goch1)
	go read(goch1,goch2)

	for{
		v,ok:=<-goch2
		if !ok{
			fmt.Println("666666666")

			break
		}
		fmt.Println("read data is",v)
	}












	//assert()
	return
	mychan = make(chan int, 4)
	fmt.Println(mychan, &mychan)
	fmt.Printf("%v , %p\n", mychan, &mychan)

	//写入数据
	mychan <- 12

	num := 66
	mychan <- num
	mychan <- num
	mychan <- num


	fmt.Println(len(mychan),cap(mychan))

	var data = <-mychan
	var data2 = <-mychan
	var data3 = <-mychan
	var data4 = <-mychan
	//var data5= <-mychan


	fmt.Println(data,data2,data3,data4,
	//,data5
	)


	fmt.Println(len(mychan),cap(mychan))




}

