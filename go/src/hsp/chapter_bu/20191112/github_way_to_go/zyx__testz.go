package main

import (
	"fmt"

	"log"
	"regexp"
	"time"
)

const (
	//TestMin = -1
	//TestA
	TestB     = 1
	TestB2    = 1
	TestB222  = 1
	TestB2221 = 1

	TestC = iota + 100
)

func main() {

	t := time.Now()
	time.Sleep(1e9)
	d := time.Now().Sub(t)

	fmt.Println(d)

	m, err := regexp.MatchString("^//(sd){1,2}open-ending$", "//sdsdopen-ending")
	fmt.Println(m, err)
	log.Println(("Listen: %d"))
	log.Println(("askuy  about douyutv  queue  design"))

	time.Sleep(1e9)
	fmt.Println("i still use")
	log.Fatal(33222)
	fmt.Println("i not use")

	//fmt.Printf("TestMin:%d\n", TestMin)
	//fmt.Printf("TestA:%d\n", TestA)
	fmt.Printf("TestB:%d\n", TestB)
	fmt.Printf("TestC:%d\n", TestC)

}
