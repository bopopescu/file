package main

import (
	"context"
	"fmt"
	"time"
	"github.com/coreos/etcd/clientv3"

	)

type asdk struct {
	Ad int
}
func Etcdt() {

	config1 := []string{"127.0.0.1:32379"}
	config := clientv3.Config{
		Endpoints:config1, // 集群地址
		DialTimeout: 5 * time.Millisecond, // 连接超时
	}

	var client *clientv3.Client
	var err error
	// 建立连接
	if client, err = clientv3.New(config); err != nil {
		fmt.Print(client,err)

	}else{
		fmt.Println(123)
	}
	fmt.Println(client,7777)
	kv := clientv3.NewKV(client)
	r1, err3 := kv.Put(context.TODO(), "get","fuckyou" )


		fmt.Println(r1,err3)

	if getResp, err := kv.Get(context.TODO(), "get", clientv3.WithPrefix()); err != nil {
		return
	}else{
		fmt.Println(getResp,6666,err)
	}
	fmt.Println(12312312)

}

func main() {

	Fav()
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
