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
	next2 := expression.Next(next)

	fmt.Println(now,"\n",next,"\n",next2)



	time.AfterFunc(next.Sub(now), func() {
		fmt.Println()
		fmt.Println("miaoji",next,"\n",time.Now())
	})


	time.Sleep(9e9)
}
