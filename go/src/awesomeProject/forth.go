package main

import (
	"strconv"
	"fmt"
	"os"
	"bufio"
)

var foo string ="dsa"
func main(){
	fmt.Println(convertobin(5))
	//printfile("abc.txt")
	}
func convertobin(n int) string{
	var result = ""
	for ;n>0;n/=2{
	  	var lsb=n%2
		fmt.Println(lsb,n,n/2)

		result= strconv.Itoa(lsb)+result
	  }
	  for a:=3;a>0;a-=1{
	  	fmt.Println(a)
	  }

	  return result
}
func printfile(filename string){
	var file,err=os.Open(filename)
	var scanner = bufio.NewScanner(file)
	for scanner.Scan(){
		fmt.Println(scanner.Text(),scanner.Bytes())
	}


	const(
		moba=iota
		dsada
		adsdasq

	)
	fmt.Println(err==nil,3333333333333)
	fmt.Println(file)





}