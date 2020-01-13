package main

import (
	"bufio"
	"fmt"
	"hsp/chapter_bu/20191217_google/u2pppw/functional/fib"
	"io"
	"strings"
	"time"
)

type intGen func() int

//实现read 接口  for Reader io.Reader
func (g intGen) Read(p []byte) (n int, err error) {
	next := g()
	if next > 100000000000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)

	// TODO incorrect if p is too small!

	time.Sleep(1e9)
	return strings.NewReader(s).Read(p)
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		fmt.Println("fuckyou", scanner.Text())
	}
}

//todo 闭包函数实现接口
func main() {
	var f intGen = fib.Fibonacci()
	printFileContents(f)
}
