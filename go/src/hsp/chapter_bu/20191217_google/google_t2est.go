package main

import (
	"encoding/json"
	"fmt"
	"hsp/chapter_bu/20191217_google/u2pppw/functional/fib"
	"math"
	"math/rand"
	"strconv"
	"testing"

	"time"

	"github.com/olivere/elastic"
	"golang.org/x/net/context"
)

var a int = 321

type fuck interface {
	Read()

}
type I interface {
	name()
}
type S struct{}

func (*S) name() {
}
func TestInterface(t *testing.T) {
	var value I = &S{}
	value.name() //可以调用

	var point = &value
	(*point).name() //不能调用
}




type DeptModeA interface {
	Name() string
	SetName(name string)
}

type DeptModeB interface {
	Relocate(building string, floor uint8)
}

type DeptModeFull interface {
	Name() string
	SetName(name string)
	Relocate(building string, floor uint8)
}

type Dept struct {
	name     string
	building string
	floor    uint8
	Key      string
}

func (self Dept) Name() string {
	return self.name
}

func (self Dept) SetName(name string) {
	self.name = name
}

func (self *Dept) Relocate(building string, floor uint8) {
	self.building = building
	self.floor = floor
}
func testInterfaceandPtr() {
	_=[]int{1}
	_=math.MaxInt32
	_=math.MinInt32
	dept1 :=
		Dept{
			name:     "MySohu",
			building: "Internet",
			floor:    7}
	switch v := (interface{}(dept1)).(type) {
	case DeptModeFull:
		fmt.Printf("The dept1 is a DeptModeFull.\n")
	case DeptModeB:
		fmt.Printf("The dept1 is a DeptModeB.\n")
	case DeptModeA:
		fmt.Printf("The dept1 is a DeptModeA.\n")
	default:
		fmt.Printf("The type of dept1 is %v\n", v)
	}

	deptPtr1 := &dept1
	if _, ok := interface{}(deptPtr1).(DeptModeFull); ok {
		fmt.Printf("The deptPtr1 is a DeptModeFull.\n")
	}
	if _, ok := interface{}(deptPtr1).(DeptModeA); ok {
		fmt.Printf("The deptPtr1 is a DeptModeA.\n")
	}
	if _, ok := interface{}(deptPtr1).(DeptModeB); ok {
		fmt.Printf("The deptPtr1 is a DeptModeB.\n")

	}
}


func  testChanSleep(fuck chan int,e chan bool){

/*
	time.Sleep(1e9)
	e <-true
	e <-true


*/
	for i:=range fuck{
		fmt.Println("模拟的=",i)
		//go func() {
			e<-false

			fmt.Println("ddddddddd")

		//}()
	}

}
func testUp(fuck chan int,e chan bool){

	go testChanSleep(fuck,e)
}
func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(
				time.Duration(rand.Intn(1500)) *
					time.Millisecond)
			out <- i
			i++

			if i>4{
				return
			}
		}
	}()
	return out
}
type Profile struct {
	Name, Height, Weight, Job, JobAddress string
	Edu, HasChild, Jiguan, Age, Marriage  string
	Sex                                   string
	Income                                string
}
func main() {
	client, i2 := elastic.NewClient(elastic.SetSniff(false))

	if i2!=nil{
		panic(i2)
	}

	fmt.Print(client,666,i2)
	item := Profile{
		Name:       "冰之泪",
		Age:        "47岁",
		Height:     "170CM",
		Marriage:   "离异",

	}


	response, i3 := client.Index().
		Index("test").
		Type("type").
		BodyJson(item).
		Do(context.Background())


	if i3!=nil{
		panic(i3)
	}
	fmt.Print(response,i3)
	fmt.Println(response.Id)
	fmt.Printf("%+v",response)
	result, i4 := client.Get().Index("test").
		Type("type").
		Id(response.Id).
		Do(context.Background())
	fmt.Println(i4)
	fmt.Printf("%+v",result)

	fmt.Println()
	fmt.Printf("%+v",string(result.Source))


	var a1 Profile
	json.Unmarshal(result.Source,&a1)
	fmt.Println()
	fmt.Println()
	fmt.Printf("%+v",a1)
	fmt.Println()

	fmt.Println(a1.Height)

	return


	j1:=make(chan int)


	go func() {
		time.Sleep(1e9)

		j1<-1234
		time.Sleep(1e9)

		//close(j1)
	}()

	for {
		fmt.Println(1111)
		result := <-j1

		fmt.Println(666,result)
	}


	return
	go func() {
		for{
			fmt.Println(<-j1)  //it is ok
		}
	}()
	fmt.Println(<-j1)   //it is deadlock

	time.Sleep(1e9)
	return
	go func() {
		fmt.Println(123)
		j1<-1
		fmt.Println(1233)
	}()
	go func() {
		fmt.Println(666)
		j1<-1
		fmt.Println(6666)
	}()

	time.Sleep(1e9)
	return  //todo  说明没有接收者也没有报错 只要是协程内
	var c1, c2 = generator(), generator()

	//var c1,c2 chan int


	//c1 = make(chan int)
	//c2= make(chan int)
	//
	//go func() {
	//	for i:=0;i<10 ;i++  {
	//
	//		c1<-i+100
	//		time.Sleep(time.Duration(rand.Intn(1500))*time.Millisecond)
	//
	//
	//	}
	//}()
	//go func() {
	//	for i:=0;i<10 ;i++  {
	//
	//		c2<-i+200
	//		time.Sleep(time.Duration(rand.Intn(1500))*time.Millisecond)
	//
	//
	//	}
	//}()

	for{
		select {
		case n1:=<-c1:
			fmt.Println("c1 out is ",n1)
		case n2:=<-c2:
			fmt.Println("c2 out is ",n2)
		default:

			//fmt.Println("receive none ")
		//return
		}
	}

	fmt.Println(123)
	return



	rrr()

	return

	//time.Sleep(1e9)

	testInterfaceandPtr()
	var a string
	a = "12"

	fuck := fib.Fibonacci()

	fmt.Println(fuck())
	fmt.Println(fuck())
	fmt.Println(fuck())
	fmt.Println(fuck())
	fmt.Println(fuck())
	fmt.Println(fuck())
	fmt.Println(fuck())
	fmt.Println(fib.Fibonacci(), 2222)

	//i, e := strconv.Atoi(a.(string))
	i, e := strconv.Atoi(a)

	fmt.Println(i, "====", e)
	//fmt.Println(strconv.Atoi(a.(string)))

}

func rrr() {
	chan1 := make(chan bool)
	fuck2 := make(chan int)
	testUp(fuck2, chan1)
	//这里形成去模拟进入fuck2的数字队列
	//go func() {
	time.Sleep(1e9)
	fuck2 <- 1235
	//fuck2<-1235
	//fuck2<-1235
	//fuck2<-1235
	//}()
	//why 定时的可以    直接 fuck2触发的超过一个都不行？  go func2 触发的又可以
	//因为协程才实现了并行
	fmt.Println(<-chan1, 555555)
	fmt.Println(<-chan1)
}

