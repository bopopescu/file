package parser

// 解析会员信息

import (
	"reflect"
	"regexp"

	"Mycrwaler/crawler-single/engine"
	"Mycrwaler/crawler-single/model"

	"github.com/rs/zerolog/log"
)

var regexs = map[string]*regexp.Regexp{
	"Age":      regexp.MustCompile(`>([^<]+)岁</div>`),
	"Img":      regexp.MustCompile(`https://photo.zastatic.com/([^<]+)\?scrop=1`),
	"Sex":      regexp.MustCompile(`扫描下载珍爱APP查看(.*)的动态`),
	"Marriage": regexp.MustCompile(`m-content-box m-des[^<]+<span [^<]+>(.*)</span></div>`),
	"Income":   regexp.MustCompile(`月收入:([^<]+)</div>`),
}
var idUrlRe = regexp.MustCompile(`http://album.zhenai.com/u/([\d]+)`)

func ParseProfile(contents []byte, url, name string) engine.ParseResult {
	rs := engine.ParseResult{}

	profile := model.Profile{Name: name}

	v := reflect.ValueOf(&profile).Elem()

	for k, r := range regexs {
		s := extractString(contents, r)
		if s != "" {
			a := v.FieldByName(k)
			if a.IsValid() {
				a.Set(reflect.ValueOf(s))
			}
		} else {
			log.Warn().Msgf("未能解析的属性：%s", k)
		}
	}

	item := engine.Item{
		Url:  url,
		Id:   extractString([]byte(url), idUrlRe),
		Type: "zhenai",

		Payload: profile,
	}
	rs.Items = []engine.Item{item}

	// 取本页面其它城市链接
	return rs
}

func extractString(c []byte, r *regexp.Regexp) string {
	match := r.FindSubmatch(c)
	if match != nil && len(match) >= 2 {
		if string(match[1]) == "他" {
			return "男"
		} else if string(match[1]) == "她" {
			return "女"

		}
		return string(match[1])
	} else {
		return ""
	}
}

type ProfileParser struct {
	userName string
	sex      string
}

func (p *ProfileParser) Parse(contents []byte, url string) engine.ParseResult {
	return ParseProfile(contents, url, p.userName)
}

func (p *ProfileParser) Serialize() (name string, args interface{}) {
	return "ProfileParser", p.userName
}

func NewProfileParser(name string) *ProfileParser {
	return &ProfileParser{
		userName: name,
	}
}
