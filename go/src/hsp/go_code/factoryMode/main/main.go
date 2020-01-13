package main

import (
	"fmt"
	"go_code/factoryMode/mode"
)

func main() {
	var stu = mode.NewStustudentdent("xss", 97.5)
	fmt.Println("", *stu)

}
