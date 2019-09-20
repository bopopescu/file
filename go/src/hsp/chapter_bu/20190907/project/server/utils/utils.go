package utils

import (
	"net"
	"fmt"
	"encoding/binary"
	"errors"
	"encoding/json"
	"hsp/chapter_bu/20190907/project/common/message"

)

//改成oop  结构体

type Transfer struct {
	Conn net.Conn
	Buf [8096]byte
}




func (this *Transfer)Readpkg() (mes message.Message, err error) {
	fmt.Println("阻塞读取中……")
	//buf := make([]byte, 1024*4)

	n, err := this.Conn.Read(this.Buf[:4]) //阻塞  前提是链接没有关闭
	if n != 4 || err != nil {
		fmt.Println("读取失败", err)
		//err=errors.New("read pkg head error")
		return
	}
	fmt.Println("读取成功，读取的buf=", this.Buf[:n])

	var pkglen uint32
	pkglen = binary.BigEndian.Uint32(this.Buf[:4])

	n2, err2 := this.Conn.Read(this.Buf[:pkglen])
	if uint32(n2) != pkglen || err2 != nil {
		fmt.Println("connread 2  wrong")

		err = errors.New("read pkg body error")

		return
	}

	//decode

	err3 := json.Unmarshal(this.Buf[:pkglen], &mes)

	if err3 != nil {
		fmt.Println("json error", err3)
		err = errors.New("read pkjson error")

		return
	}

	return

}

func (this *Transfer)WritePkg( data []byte) (err error) {

	var pkgLen uint32

	//this is 76 124 likes

	pkgLen = uint32(len(data))

	//var bytes [5]byte
	var bytes []byte = make([]byte, 4)

	//将数字变成 字节切片 todo
	binary.BigEndian.PutUint32(bytes[0:4], pkgLen)

	n, err := this.Conn.Write(bytes)
	if n != 4 || err != nil {
		fmt.Println("发送失败", err)
		return
	}
	fmt.Println("服务端发送信息长度成功", len(data), data, string(data))

	ne, err := this.Conn.Write(data)
	fmt.Println("send msg length ", ne)
	if ne != int(pkgLen) || err != nil {
		fmt.Println("发送body失败", err)
		return
	}

	return
}