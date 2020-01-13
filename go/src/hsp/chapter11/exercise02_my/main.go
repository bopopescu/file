package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type Hero struct {
	Name string
	Age  int
}

type Heroslice []Hero

func (this Heroslice) Len() int {
	return len(this)
}

func (this Heroslice) Less(i, j int) bool {

	return this[i].Age < this[j].Age
}

func (this Heroslice) Swap(i, j int) {
	//temp:=this[i]
	//this[i]=this[j]
	//this[j]=temp

	this[i], this[j] = this[j], this[i]
}

func main() {
	var slicearr = []int{1, 3, 33, 123, 99}

	sort.Ints(slicearr)

	fmt.Println(slicearr)

	var heros Heroslice
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		hero := Hero{
			Name: fmt.Sprintf("hero~%d", rand.Intn(100)),
			Age:  rand.Intn(30),
		}
		heros = append(heros, hero)
	}

	fmt.Println(heros)

	sort.Sort(heros)
	for k, v := range heros {
		fmt.Println(v, k)
	}
}
