package main

import (
	"fmt"

	"github.com/mediocregopher/radix.v2/pool"
)

func main() {
	redisPool, err := pool.New( "tcp", "localhost:6379", 5 );

	cmd := redisPool.Cmd("PING")
	ret:= redisPool.Cmd( "set", "fuck", "you1" )

	fmt.Print(redisPool,err,cmd,ret)


}
