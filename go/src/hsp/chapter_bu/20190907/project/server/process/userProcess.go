package process

import (
	"net"
	"encoding/json"
	"fmt"
	"hsp/chapter_bu/20190907/project/common/message"
	"hsp/chapter_bu/20190907/project/server/utils"
)

type UserProcess struct {
	Conn net.Conn

}
func (this *UserProcess)ServerProcessLogin( mes *message.Message) (err error) {
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)

	if err != nil {
		fmt.Println("json un errr", err)
		return
	}

	var resMes message.Message
	resMes.Type = message.Loginmesrestype

	var loginresmes message.LoginMesRes

	if loginMes.Userid == 100 && loginMes.Userpwd == "123456" {

		loginresmes.Code = 200

	} else {

		loginresmes.Code = 500
		loginresmes.Error = "用户名密码错误"

	}

	data, err := json.Marshal(loginresmes)
	if err != nil {
		fmt.Println("json encode res", err)
		return
	}

	resMes.Data = string(data)

	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json encode res", err)
		return
	}

	tf:= &utils.Transfer{
		Conn:this.Conn,
	}
	err = tf.WritePkg( data)
	return

}
