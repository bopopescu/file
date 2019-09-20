package main

import (
	"fmt"
	"time"
)

















func raw(rawdata chan int){


	for i:=1;i<=8000;i++{
		rawdata<- i
	}
	close(rawdata)

}


func calprime(rawdata chan int,dealdata chan int,endchan chan bool){

	for{
		v,ok:=<-rawdata
		if(ok){
			if v%2==0{
				dealdata<-v

			}
		}else{
			break
		}
	}

	endchan<-true

	//if(len(endchan)==4){
	//	close(endchan)
	//}


}


func main01(){
	btime:=time.Now().UnixNano()

	rawdata:= make(chan int ,8000)
	dealdata:=make(chan int,5000)

	for i:=1;i<=8000;i++{
		rawdata<- i
	}

	close(rawdata)
	for{
		v,ok:=<-rawdata
		if(ok){
			if v%2==0{
				dealdata<-v

			}
		}else{
			break
		}
	}
	close(dealdata)

	//for{

			fmt.Println("asdasdasd",dealdata)
			//for v:=range dealdata{
			//	fmt.Println(v)
			//}
			//break

	//}
	etime:=time.Now().UnixNano()
	fmt.Println(etime-btime)

}


func main0(){
	btime:=time.Now().UnixNano()
	rawdata:= make(chan int ,2000)

	dealdata:=make(chan int,5000)

	endchan :=make(chan bool,4)
	go raw(rawdata)
	go calprime(rawdata,dealdata,endchan)
	go calprime(rawdata,dealdata,endchan)
	go calprime(rawdata,dealdata,endchan)
	go calprime(rawdata,dealdata,endchan)




	for{
		abc := len(endchan)
		if abc==4{
			close(dealdata)
			fmt.Println("asdasdasd",dealdata)
			//for v:=range dealdata{
			//	fmt.Println(v)
			//}
			fmt.Println(123)
			break
		}
	}

	etime:=time.Now().UnixNano()

	fmt.Println(etime-btime)
}
