package main

import (
	"net"
	"fmt"
	"hsp/chapter_bu/20190907/project/common/message"
	"hsp/chapter_bu/20190907/project/server/process"
	"io"
	"hsp/chapter_bu/20190907/project/server/utils"
)


type Processor struct {

	Conn net.Conn

}
//主控代码
func (this *Processor)ServerProcessMes( mes *message.Message) (err error) {
	switch mes.Type {
	case message.Loginmestype:


		up :=process.UserProcess{
			Conn:this.Conn,
		}
		err = up.ServerProcessLogin( mes)



	case message.RegMesType:

	default:
		fmt.Println("type not exxist")

	}
	return
}




//处理客户端的通讯
func (this *Processor)processmr()(err error) {
	//循环读客户端发送的信息


	for {

		tf:= &utils.Transfer{
			Conn:this.Conn,
		}
		mes, err := tf.Readpkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端已经退出", this.Conn.RemoteAddr(), err) //todo   notice here  EOF.
				return err

			} else {
				fmt.Println("goprocess:read wrong", err)

				return err
			}

		}


		this.ServerProcessMes( &mes)

		fmt.Println("mes=", mes)
		//校验 然后开个协程？
	}

	return err

}