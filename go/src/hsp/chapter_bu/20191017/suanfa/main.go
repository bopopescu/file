package main

import (
	"errors"
	"fmt"
	"os"
)

type Queue struct {
	maxSize int
	array   [5]int
	front   int
	rear    int
}

func main() {
	fmt.Println("this is suanfa ")

	queue := &Queue{
		maxSize: 5,
		array:   [5]int{},
		front:   -1,
		rear:    -1,
	}
	fmt.Println(queue.showYouFive())
	fmt.Println(queue.showYouFive2())

	var key string

	var val int
	for {
		fmt.Println("1=>push")

		fmt.Println("2=>pop")

		fmt.Println("3=>showqueue")

		fmt.Scanln(&key)

		fmt.Println("your key is ", key)

		switch key {
		case "1":
			n, err2 := fmt.Scanln(&val)
			fmt.Println(n, err2)
			fmt.Println("you type in is ", val)
			err := queue.addqueue(val)
			if err != nil {
				fmt.Println("insert wrong", err, 99999999)
			}
			fmt.Println(err, 888888)

		case "2":
			i, err := queue.getqueue()
			if err != nil {
				fmt.Println("get wrong", err)
			}
			fmt.Println(i)

		case "3":

			queue.showqueue()
		case "4":
			os.Exit(0)
		default:

			fmt.Println("type wrong")

		}
	}
}

func (this *Queue) addqueue(val int) (err error) {

	if this.rear == this.maxSize-1 {

		return errors.New("quque full")
	}

	this.rear++

	this.array[this.rear] = val
	//15171505290

	fmt.Println(val)
	return
}

func (this *Queue) showqueue() {

	fmt.Println()

	fmt.Println("let us show ")

	for i := this.front + 1; i <= this.rear; i++ {
		fmt.Println(i, this.array[i])
	}

	fmt.Println()
}

func (this *Queue) getqueue() (val int, err error) {

	if this.front == this.rear {
		return 0, errors.New("empty queue")
	}
	this.front++

	return this.array[this.front], nil
}

//hash  list set sadd  kv

func (this *Queue) showYouFive() []int {

	var arr []int = make([]int, 6)
	arr[0] = 12
	arr[1] = 324
	arr[2] = 322
	arr[3] = 323
	arr[4] = 321
	arr[5] = 321

	demo(arr)
	return arr
}

func demo(demo []int) {
	(demo)[0] = 33332
}

func (this *Queue) showYouFive2() [3]int {

	var arr [3]int
	arr[0] = 12
	arr[1] = 324
	arr[2] = 322

	demo2(&arr)
	return arr
}

func demo2(demo *[3]int) {
	(*demo)[0] = 33332
}
