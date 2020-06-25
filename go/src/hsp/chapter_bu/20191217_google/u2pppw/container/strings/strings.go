package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Yes我爱慕课网!" // UTF-8
	fmt.Println(s)

	for _, b := range []byte(s) {
		fmt.Printf("%X ", b)
	}
	fmt.Println()

	for i, ch := range s { // ch is a rune
		fmt.Printf("(%d -- %X) ", i, ch)
	}
	fmt.Println()

	fmt.Println("Rune count:",
		utf8.RuneCountInString(s))

	bytes := []byte(s)
	fmt.Println("bts", bytes)
	for len(bytes) > 0 {
		fmt.Println()

		fmt.Println("lenis", len(bytes))
		ch, size := utf8.DecodeRune(bytes)
		fmt.Println(ch, "----", size)
		bytes = bytes[size:]
		fmt.Printf("%c 】】】】】", ch)
		fmt.Println()
	}
	fmt.Println()

	for i, ch := range []rune(s) {
		fmt.Printf("(%d %c) ", i, ch)
	}

	fmt.Printf("( %c) ", 32596)
	fmt.Printf("( %c) ", 32597)
	fmt.Printf("( %c) ", 32598)
	fmt.Printf("( %c) ", 32599)
	fmt.Printf("( %c) ", 32600)

	fmt.Println()
}
