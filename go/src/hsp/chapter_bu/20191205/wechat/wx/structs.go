package wx

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"time"
)

type Base struct {
	FromUserName CDATAText
	ToUserName   CDATAText
	MsgType      CDATAText
	CreateTime   CDATAText
}

func (this *Base) testFuntion() {

}

type vable struct {
	faker int
	BETTY string
}

func (test *Base) isExisted() {

}
func (b *Base) InitBaseData(w *WeixinClient, msgtype string) {

	b.FromUserName = value2CDATA(w.Message["ToUserName"].(string))
	b.ToUserName = value2CDATA(w.Message["FromUserName"].(string))
	b.CreateTime = value2CDATA(strconv.FormatInt(time.Now().Unix(), 10))
	b.MsgType = value2CDATA(msgtype)
}

type CDATAText struct {
	Text string `xml:",innerxml"`
}

func test() {
	map1 := map[int]string{0: "asd", 1: "oopp"}

	map1[2] = "fuckyou"
	map1[100] = "fuckyouone hundred times"
	fmt.Printf("the p is %p ,  the x is %x", map1, map1)
	fmt.Println(map1)
}

type TextMessage struct {
	XMLName xml.Name `xml:"xml"`
	Base
	Content               CDATAText
	whatifIdontExpthisVar map[string]interface{}
}
