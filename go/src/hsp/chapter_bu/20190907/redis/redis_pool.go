package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

// redis  连接池\

var pool *redis.Pool

func init() {

	pool = &redis.Pool{
		MaxIdle:     8,
		MaxActive:   0,
		IdleTimeout: 100,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "127.0.0.1:6379")
		},
	}
}

func main() {

	conn := pool.Get()
	fmt.Println(conn)
	fmt.Printf("%T,%v", conn, conn)
	defer conn.Close() //这里关闭 会放回连接池

	ji, err := redis.Int(conn.Do("get", "a1"))

	fmt.Println(ji, err)

}
