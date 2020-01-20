package main

import (
	"fmt"
	"time"

	"github.com/gorhill/cronexpr"
)

func main() {






	expression, e := cronexpr.Parse("*/4 * * * * * *")
	fmt.Print(e)


	fmt.Println()
	now:=time.Now()
	next := expression.Next(now)
	fmt.Println(now,next)



	time.AfterFunc(next.Sub(now), func() {
		fmt.Println("miaoji",next)
	})


	time.Sleep(9e9)
}
