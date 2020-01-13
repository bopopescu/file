package main

import (
	"fmt"
	"strings"
	"time"
)

func askuy() {
	var a [3]string
	a[0] = "asd"
	a[1] = "dddd"
	a[2] = "ddddxcasd"
	fmt.Println(a, a[:2])

	var aa = []string{"dd", "ff"}
	fmt.Println(aa, aa[:2])

	var data = make(map[string]string)
	data["test1"] = "dsadas"
	data["test2"] = "dsdadwqeqe"
	fmt.Println(data)

	var mixdata = []interface{}{"football", 12333}
	fmt.Println(mixdata)

	var mixed interface{}
	mixed = 123

	mixed = []string{"A", "b"}
	fmt.Println(mixed)

	fmt.Println(time.Now().Format("2018-10-10 15:00:00"))

	str := "f,a,k,e,r"
	array := strings.Split(str, ",")
	fmt.Println(str, array, array[0], array[1:])

	var imp = map[string]string{
		"us": "dsd",
		"cn": "lpl"}
	fmt.Println(imp["23"], imp["us"])

	var vzx string

	vzx, exists := imp["us"]

	if !exists {
		fmt.Println("no", vzx, "infinity")
	} else {
		fmt.Println("yesyes", vzx, "meceydes")
	}
}

func main() {

	askuy()
	fmt.Println(test())

	var i int = 123

	var a = []byte("sdasd")

	var r = []byte{4, 3}

	var arr = [4]int{1, 2, 3, 4}

	var arr2 = [3][2]int{[2]int{1, 2}, [2]int{3, 4}, [2]int{5, 6}}
	fmt.Println(i, a, r, arr)
	var arr3 = []int{3, 4, 6}
	s3 := arr3[1:3]

	var array = []int{1, 2, 3, 4}

	array2 := append(array, 32)
	array[1] = 123
	fmt.Println(array, array2)
	fmt.Println(s3)

	fmt.Println(arr2, arr3)
	for o, p := range arr2 {
		fmt.Println(o, p)
	}

	switch i {
	case 1233, 123:
	case 223:
		fmt.Println(2132)
	case 2232:
		fmt.Println(21323)
	default:
		fmt.Println("this is default")
	}

	list := []string{"a", "b", "c", "d"}

	for k, v := range list {
		fmt.Println(k, v)
	}

	fmt.Println("\n")
	test1()
}

func test1() {

	c := "test1"
	c2 := test1

	fmt.Println(c, c2)
	for a, b := range "jacker" {
		fmt.Println(a, string(b))
	}
}

func test() string {
	return "123"
}
