package fib

// 1, 1, 2, 3, 5, 8, 13, ...
func Fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b //1   1     1    2      2   3    3  5   5
		return a
	}
}

func iwritefib() func () int{

	a,b:=0,1

	return func()int{
		a,b=b,a+b
		return a
	}


	//1 1    1
	// 1  2    1
	 // 2  3   2
	 // 3  5   3
	 //  5   8   5
	 //  8   13  8
	 //13   21   13

}


func iwrite_easy_understand() func() int{
	a, b := 1, 1
	return func() int {
		c:=a
		a=b

		b=a+c
		return c
	}

	// 1 c=1  a=1  b=2
	// 1 c=1  a=2  b=3
	// 2 c=2  a=3  b=5
	// 3 c=3  a=5  b=8
	// 5 c=5  a=8  b=13
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
