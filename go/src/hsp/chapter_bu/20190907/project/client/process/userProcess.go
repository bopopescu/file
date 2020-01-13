package process

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"hsp/chapter_bu/20190907/project/client/utils"
	"hsp/chapter_bu/20190907/project/common/message"
	"net"
	"time"
)

type UserProcess struct {
}

func (this *UserProcess) Login(userid int, userpwd string) (err error) {

	fmt.Println("your type is ", userid, "----", userpwd)

	//连接服务器

	conn, err := net.Dial("tcp", "127.0.0.1:8889")
	tf := &utils.Transfer{
		Conn: conn,
	}

	if err != nil {
		fmt.Println("dial wrong", err)
	}

	var mes message.Message
	mes.Type = message.Loginmestype

	var loginMes message.LoginMes
	loginMes.Userid = userid
	loginMes.Userpwd = userpwd
	tmp, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json wrong", err)
		return
	}

	mes.Data = string(tmp)

	data2, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("json2 wrong", err)
	}

	//先发送长度

	var pkgLen uint32

	//this is 76 124 likes

	pkgLen = uint32(len(data2))

	//var bytes [5]byte
	var bytes []byte = make([]byte, 4)

	//将数字变成 字节切片 todo
	binary.BigEndian.PutUint32(bytes[0:4], pkgLen)

	n, err := conn.Write(bytes)
	if n != 4 || err != nil {
		fmt.Println("发送失败", err)
		return
	}

	fmt.Println("客户端发送信息长度成功", len(data2), data2, string(data2))

	n2, err := conn.Write(data2)
	fmt.Println("send msg length ", n2)
	if err != nil {
		fmt.Println("发送body失败", err)
		return
	}

	//处理server返回的消息

	mes2, err := tf.Readpkg()
	if err != nil {
		fmt.Println("接收服务器消息失败", err)
		return
	}

	var loginResmesp message.LoginMesRes
	_ = json.Unmarshal([]byte(mes2.Data), &loginResmesp)

	if loginResmesp.Code == 200 {

		fmt.Println("login succ")

		go serverProcessmes(conn)

		for {
			ShowMenu()
		}

	} else {
		fmt.Println("login fail  ,account not correct")

	}

	time.Sleep(time.Second * 10) //just for test

	defer conn.Close()

	return nil
}
