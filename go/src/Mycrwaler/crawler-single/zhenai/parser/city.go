package parser

import (
	"regexp"

"Mycrwaler/crawler-single/engine"

)

var (
	profileRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/\w+)"[^>]*>([^<]+)</a>`)

	cityUrlRe = regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/wuhan/[^"]+)`)
	//cityUrlRe = regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/wuhan/\d+)"`)
)

func ParseCity(contents []byte, _ string) engine.ParseResult {
	rs := engine.ParseResult{}

	match := profileRe.FindAllSubmatch(contents, -1)
	limit:=100
	for _, m := range match {


		//fmt.Println(string(m[1]),string(m[2]),"--人")


		rs.Requests = append(rs.Requests, engine.Request{
			Url:   string(m[1]),//继续爬取   个人主页
			Parse: NewProfileParser(string(m[2])),
		})

	}

	// 取本页面其它城市链接
	match = cityUrlRe.FindAllSubmatch(contents, -1)
	for _, m := range match {
		//fmt.Println(string(m[1]),"--城")

		rs.Requests = append(rs.Requests, engine.Request{
			Url:   string(m[1]),
			Parse: engine.NewFuncParser(ParseCity, "ParseCity"),
		})
		limit--
		if limit<=0{
			break
		}
	}

	return rs
}
