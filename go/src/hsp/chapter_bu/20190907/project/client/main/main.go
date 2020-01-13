package main

import (
	"fmt"
	"hsp/chapter_bu/20190907/project/client/process"
	"os"
)

var userid int
var userpwd string

func main() {
	var key int

	var loop = true

	for loop {
		fmt.Println("--you are login--")
		fmt.Println("\t\t\t 1 login")
		fmt.Println("\t\t\t 2 reg")
		fmt.Println("\t\t\t 3 exit")
		fmt.Println("\t\t\t plz choose ")

		fmt.Scanf("%d\n", &key)
		//fmt.Scanln(&key)

		//fmt.Println(key)
		switch key {
		case 1:

			fmt.Println("print your uid")

			fmt.Scanf("%d\n", &userid)

			fmt.Println("print your passwword")

			fmt.Scanf("%s\n", &userpwd)

			up := &process.UserProcess{}
			up.Login(userid, userpwd)

			//_=login(userid,userpwd)
			loop = false
			fmt.Println("logindo")
			break

		case 2:
			fmt.Println("regdo")
			loop = false

			break
		case 3:
			fmt.Println("exitdo")
			loop = false

			fmt.Println("are you ok")
			os.Exit(0)
			break

		default:

			fmt.Println("type wrong", key)
		}

		if !loop {

			//break
		}
	}

}
