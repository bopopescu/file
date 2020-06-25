package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
	"sync"

)

func main() {
	var(
		cmd *exec.Cmd
		err error
		output []byte

	)

	files:=make([]string,0,10)
	//dir,err:=ioutil.ReadDir("/Users/infy/Desktop/file/go")
	//for _,dr :=range dir{
	//	if dr.IsDir(){
	//		files=append(files,dr.Name())
	//	}
	//}
	files = []string{
		"tsuc","ask","news","pay","config","service","tingsong","tszfbserver",
		"PayCenter","upload","docs","master","points",
	}


	//fmt.Println(files, len(files))
	//return
	//  theroutes :=[]string{"/Users/infy/Desktop/file/go/pkg","/Users/infy/Desktop/file/go/bin"}



	cmd=exec.Command("/bin/bash","-c","sudo echo $USER")
	fmt.Print(err,cmd)
	if output, err = cmd.CombinedOutput();err!=nil{
		panic(err)
	}
	fmt.Println()
	user := string(output)
	user = strings.Replace(user, "\n", "", -1)
	fmt.Println(string(output))
	devspace :="/home/"+user+"/devspace/"

	fmt.Println("kaishi")


	lock :=sync.Mutex{}



	wg:=sync.WaitGroup{}
	wg.Add(len(files))

	  for i,v:=range files{

	  	go func(s string) {
			var cmd *exec.Cmd

			var outputs []byte
			//if s=="src"{
			//	fmt.Println("aaaaaa")
			//	s = s + ";sleep 2"
			//}

			s = devspace+s
			//fmt.Println(s,6666)
			//cmd=exec.Command("/bin/zsh","-c"," cd "+ s +";sudo  ls -al;pwd")
			lock.Lock()
			cmd=exec.Command("/bin/bash","-c"," cd "+s+";sudo pwd;zp start")
			//err=cmd.Run()
			fmt.Println(err,cmd)


			var out bytes.Buffer

			var stderr bytes.Buffer

			if outputs, err = cmd.CombinedOutput();err!=nil{
				fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
				fmt.Println("Result: " + out.String())

				//panic(err)
			}


			fmt.Println()
			fmt.Println(string(outputs))
			wg.Done()
			lock.Unlock()

		}(v)


	  	fmt.Println(i,v)
	  }


	  wg.Wait()

	  fmt.Println("done all")

	  return

}
