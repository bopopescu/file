package main

import "fmt"

func main() {
	var a2 = new([5]int)

	//var p [5]int = new([5]int)
	var p *[]int = new([]int)

	fmt.Println(p, *p, p == nil, *p == nil, p == new([]int))

	fmt.Println(777)
	fmt.Printf("%p-----%p\n", p, *p)
	fmt.Println(555)
	var a1 [5]int

	var a3 = newInt()
	fmt.Println(a2, a1, a3)
	fmt.Println(*a2 == a1)

	var jmap = make(map[int]string, 2)
	var jchan = make(chan int, 5)
	var jslice = make([]int, 2, 3)

	jchan <- 123
	jslice[1] = 1
	jslice[0] = 13

	jslice = append(jslice, 14)
	jslice = append(jslice, 414)
	jslice = append(jslice, 414)
	//jslice[23]=13
	jmap[1] = "sda"
	jmap[12] = "sda"
	jmap[121] = "sda"
	jmap[122] = "sda"
	jmap[123] = "sda"
	jmap[1232] = "sda"
	fmt.Println(jmap, len(jmap))
	fmt.Println(jslice, len(jslice), cap(jslice))

	fmt.Println((*a2)[1], &a2[2], &a2[3], &a2[4])

	a := [5]int32{1, 2, 3, 4, 5}
	b := a
	c := b
	fmt.Printf("数组 a - 值：%v，指针：%p\n", a, &a)
	fmt.Printf("数组 b - 值：%v，指针：%p\n", b, &b)

	fmt.Println(&a, &a[0], &a[1], &a[2], &a[3], &a[4])
	fmt.Println(&b, &b[0], &b[1], &b[2], &b[3], &b[4])
	fmt.Println(&c, &c[0], &c[1], &c[2], &c[3], &c[4])

	var auy map[string]string
	//在使用map前，需要先make , make的作用就是给map分配数据空间
	auy = make(map[string]string, 10)
	auy["no1"] = "宋江" //ok?
	auy["no2"] = "吴用" //ok?
	auy["no1"] = "武松" //ok?
	auy["no3"] = "吴用" //ok?
	fmt.Println(auy)
}

func newInt() *[5]int {
	var i [5]int
	return &i
}
