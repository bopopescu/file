package parser

import (
	"thehappymouse/ccmouse/crawler/engine"
	"fmt"
	"regexp"
)

const cityListRegex = `<a href="(http://www.zhenai.com/zhenghun/\w+)"[^>]*>([^<]+)</a>`

func ParseCityList(contents []byte, _ string) engine.ParseResult {
	rs := engine.ParseResult{}

	reg := regexp.MustCompile(cityListRegex)
	match := reg.FindAllSubmatch(contents, -1)
	limit:=10
	for _, m := range match {

//没有return items


//只有具体city才 items

		fmt.Println(string(m[1]))
		rs.Requests = append(rs.Requests, engine.Request{
			Url:   string(m[1]),//理解为继续爬取 城市页
			Parse: engine.NewFuncParser(ParseCity, "ParseCity"),
		})
		limit--
		if limit<=0{
			break
		}
	}

	return rs
}
