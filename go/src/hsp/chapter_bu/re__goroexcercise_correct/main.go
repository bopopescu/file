package main

import (
	"fmt"
	"time"
)

func put(intchan chan int){

	for i:=1;i<=130000;i++{
		intchan	<-i
	}
	close(intchan)
}

func prime(intchan chan int,primechan chan int,exitchan chan int64){
	//var num int
	var flag bool
	for{
		num,ok:=<-intchan
		if(!ok){
			break
		}
		flag=true
		for i:=2;i<num;i++{
			if num%i==0{
				flag=false
				break
			}
		}

		if flag{
			primechan<-num

		}

	}
	fmt.Println("one goroutine done")
	exitchan<- (time.Now().UnixNano())
}

func main(){
	btime:=time.Now().Unix()
//
//	for num:=1;num<=130000;num++{
//		for i:=2;i<num;i++{
//			if num%i==0{
//				break
//			}
//		}
//	}
//	etime2:=time.Now().Unix()
//	fmt.Println(etime2-btime)
//return

	intchan:=make(chan int,80000)
	primechan:=make(chan int,200000)
	//exitchan:=make(chan bool,16)
	exitchan:=make(chan int64,16)
	//exitchan2:=make(chan int64,16)

	go put(intchan)
	var gonum =4   //其实上面这个go 也占一个协程  所以估计 3个和4个下列协程应该效率一致
	for i:=1;i<=gonum;i++{
		go prime(intchan,primechan,exitchan)

	}



	//用匿名协程跑这个
	//go func(){

		fmt.Println(time.Now().UnixNano())
//close(exitchan2)
	//todo  这里只能为4   携程数量  因为没有关闭exitchan   所以只能读4个  这里应该是阻塞读

		for i:=1;i<=gonum;i++{
			fmt.Println("\n in=>",time.Now().UnixNano())

			v:=<-exitchan

			// todo  为什么这里用这个没有写入管道的exitchan2 会直接超时  ===》 底层底层 还是底层
			//<-exitchan2


			fmt.Printf("%T %v\n",v,v)
			fmt.Println("\n out=>",time.Now().UnixNano())
		}
		close(primechan)
		etime:=time.Now().Unix()
		fmt.Println(etime-btime)

	//}()


	//var times = 0
	//for{
	//	//如果不关闭   最后一次没有打印出ok 为false的
	//	times++
	//	sd,ok:=<-primechan
	//	fmt.Println(ok,times,"value",sd)
	//
	//
	//	if !ok{
	//		break
	//	}
	//	//fmt.Println("sushu,",res)
	//}

}