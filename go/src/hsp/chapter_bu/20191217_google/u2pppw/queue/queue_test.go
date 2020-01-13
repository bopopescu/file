package queue

import (
	"fmt"
	"os"
	"testing"
)

func ExampleQueue_Pop() {
	q := Queue{1}
	q.Push(2)
	q.Push(3)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())

	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	str, _ := os.Getwd()
	fmt.Println(str)
	// Output:
	// 1
	// 2
	// false
	// 3
	// true
	///Users/infy/Desktop/file/go/src/hsp/chapter_bu/20191217_google/u2pppw/queue
}

func TestHello(t *testing.T) {
	t.Logf("执行正确hello")
}

func TestOc(t *testing.T){
	t.Logf("执行正确hello")

}