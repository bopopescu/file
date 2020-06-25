package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {


	f:=flag.String("file","/tmp/acc.log","help type filename")

	flag.Parse()

	file, e := os.OpenFile(*f, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)

	log.SetOutput(file)

	fmt.Print(file,e,7666,*f)

	log.Println("fuck")
	file.Write([]byte("askuy\naskuy\naskuy\naskuy\n"))
}
