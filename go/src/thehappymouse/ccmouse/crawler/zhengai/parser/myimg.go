package parser

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/thehappymouse/ccmouse/crawler/engine"
)

var (
	//imgRexp2 = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/\w+)"[^>]*>([^<]+)</a>`)
	imgRexp2 = regexp.MustCompile(`<img src="([^<]+)" .* width="`)

	download = regexp.MustCompile(`http://cdn.clearloveuzi.cn/(.*)/(.*)`)
	)

func Parseimgy(contents []byte, _ string) engine.ParseResult {
	rs := engine.ParseResult{}
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	match := imgRexp2.FindAllSubmatch(contents, -1)
	for _, m := range match {

		resp, _ := http.Get(string(m[1]))
		body, _ := ioutil.ReadAll(resp.Body)

		thename:=""
		dl := download.FindAllSubmatch(m[1], -1)
		for _, ol := range dl {
			if  len(ol)>=2{
				thename = string(ol[2])
			}else{
				thename=""
				continue
			}
		}


		out, _ := os.Create(thename)

		io.Copy(out, bytes.NewReader(body))

		log.Info().Msg(string(m[1]))


		//爬取后下载
	}


	return rs
}
