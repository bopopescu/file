package fib

// 1, 1, 2, 3, 5, 8, 13, ...
func Fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b //1   1     1    2      2   3    3  5   5
		return a
	}
}

func fibHsp() func() int {

	a, b := 1, 1
	return func() int {
		c := a
		a = b
		b = a + c
		return c
	}

}

//1  1
