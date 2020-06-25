package main

import (
	"fmt"
	"os"

	"bufio"

	"hsp/chapter_bu/20191217_google/u2pppw/functional/fib"
)

func tryDefer() {
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
		if i == 5 {

			fmt.Println(123)

			//panic("offuoi")

			// Uncomment p anic to see
			// how it works with defer
			//panic("printed too many")
		}
	}
}

func writeFile(filename string) {
	//file, err := os.Create(filename)

	// If there is an error, it will be of type *PathError.
	file, err := os.OpenFile(filename,
		os.O_EXCL|os.O_CREATE|os.O_WRONLY, 0777)

	//err = errors.New("fuck fuck you bitch ")  //not a *os.PathError
	if err != nil {
		fmt.Println(err.Error(), "err.error()", err)

		if pathError, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			fmt.Printf("fuck %s\n, %s\n, %s\n",
				pathError.Op,
				pathError.Path,
				pathError.Err)

			fmt.Println("test open ",err,ok)
		}
		//return
	}



	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	f := fib.Fibonacci()

	for i := 0; i < 10; i++ {
		fmt.Fprintln(writer, f())
	}
}

func main() {
	tryDefer()

	fmt.Println(666)
	writeFile("fib231.txt")
	fmt.Println(777)

}
