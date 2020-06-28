package main

import (
	"fmt"
	"time"
	)

type asdk struct {
	Ad int
}


func main() {
	ji :=make(chan int)
	fmt.Println(123)

	ji<-111

	fmt.Println(<-ji)

	fmt.Println(123)
	return

	var log1 *asdk
	log1=&asdk{}
	fmt.Println(log1,&log1)
	log1=&asdk{Ad:123}
	fmt.Println(log1,&log1)
	log1 = nil
	fmt.Println(log1,&log1)
	return
	// 调度的延迟定时器
	//scheduleTimer := time.NewTimer(4e9)

	intchan :=make(chan int)
	intchan2 :=make(chan int)
	go func() {
		time.Sleep(3e9)
		//
		//intchan<-1
		//intchan2<-1
	}()

	fmt.Println(222)

	for{


	select {


		case jobEvent := <- intchan:	//监听任务变化事件

			fmt.Println(jobEvent)
		case jobEvent2 := <- intchan2:	//监听任务变化事件

		fmt.Println(jobEvent2)

	default:
		fmt.Println(1)
		//case <- scheduleTimer.C:	// 最近的任务到期了
			//default:
			//	fmt.Println(333)
	}

	}
	fmt.Println(123)
}
