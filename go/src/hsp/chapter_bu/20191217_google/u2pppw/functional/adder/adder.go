package main

import (
	"fmt"
)

func adder() func(int) string {
	sum := 0
	return func(v int) string {
		sum += v
		s := fmt.Sprintf("0+1+...+%d = %d", v, sum)
		return s
	}
}

//0 1  3  6 10 15 21 28 36 45

type iAdder func(int) (int, iAdder)

func adder2(base int) iAdder {
	return func(v int) (int, iAdder) {
		return base + v, adder2(base + v)
	}
}

func main() {

	//testAdder()
	//
	//return

	// a := adder() is trivial and also works.
	a := adder2(0)
	for i := 0; i < 10; i++ {
		var s int
		s, a = a(i)
		fmt.Printf("0 + 1 + ... + %d = %d\n",
			i, s)
	}
}

func testAdder() {
	b := adder()
	for i := 0; i < 10; i++ {
		fmt.Println(b(i))
	}
}
