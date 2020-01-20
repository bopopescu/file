package main

import (
	"fmt"
	"os/exec"
	"time"

	"golang.org/x/net/context"
)


type Resultcmd struct {

	err error
	output string
}

func main() {
	var(
		cmd *exec.Cmd
		err error
		output []byte
		ctx  context.Context
		cancel context.CancelFunc
		cmdchan chan Resultcmd
	)
/*
	cmd=exec.Command("/bin/zsh","-c","sleep 2;echo 2;ls -al;pwd")
	//err=cmd.Run()
	fmt.Print(err,cmd)
	if output, err = cmd.CombinedOutput();err!=nil{
		panic(err)
	}
	fmt.Println()
	fmt.Println(string(output))*/




	cmdchan=make(chan Resultcmd)


	ctx, cancel = context.WithCancel(context.TODO())


	go func() {
		cmd = exec.CommandContext(ctx, "/bin/zsh", "-c", "sleep 2;echo 2;ls -al;pwd")

		output, err = cmd.CombinedOutput()

		cmdchan<-Resultcmd{
			err:err,
			output:string(output),
		}

		fmt.Println(333,err)


	}()



	time.Sleep(1e9)
	cancel()

	res ,ok :=<-cmdchan


	fmt.Println(ok)
	fmt.Printf("%+v",res)

	}
