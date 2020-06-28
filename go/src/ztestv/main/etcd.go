package main

import (
	"fmt"
	"time"
	"github.com/coreos/etcd/clientv3"
	"context"
)

//Todo 为什么没有报main重复的错误呢
func main() {
	Etcdt()
	fmt.Print(123)

}

func Etcdt(){

	fmt.Println(3)
}
func Fav() {

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
