package main

import (
	"fmt"
	"io/ioutil"
	"math"
)

func grade(score int) string {
	var g = ""
	switch {
	case score < 60:
		g = "f"
	case score < 80:
		g = "a"
	case score > 80:
		g = "s"
	default:
		panic(fmt.Sprintf("wrong score:%d", score))
	}
	return g
}
func main() {
	const ask string = "dsd"
	const uy int = 123
	const (
		tes  = "as"
		e, d = 5, 6
	)

	var gg = grade(81)
	fmt.Println(gg)
	const a, b = 3, 4

	var c int
	c = int(math.Sqrt(a*a + b*b))

	fmt.Println(c, e, d)

	array()
	file()
}

func array() {
	const (
		cpp  = iota
		java = 2
		php  = 3
		js   = 4
	)

	const (
		moba = iota
		dsada
		adsdasq
	)
	fmt.Println(cpp, java, php, dsada)
}

func file() {
	const filename = "abc.txt"
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s"+"\n", contents)

		fmt.Println(contents, 12312)
	}
}
