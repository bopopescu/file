package main

import (
	"fmt"
	"net/http"

	"bufio"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/net/html/charset"
	"io/ioutil"
	"regexp"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog"
	"os"
	"golang.org/x/text/transform"
)

func main(){

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})


	log.Print("hello world")

	log.Info().Msg("Got Item: $%d %2")
	log.Warn().Msg("dsa")

	//resp, err := http.Get("http://www.zhenai.com/zhenghun")
	resp, err := http.Get("https://www.clearloveuzi.cn/2020/04/11/144.html")

	if err!=nil{
		panic(err)
	}

	fmt.Println(resp,"hey eva of 21century",err)


	fmt.Println("aaaaaaaaddddddddddaaaaaaaaaaaaaa")
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK{

		//response, i := httputil.DumpResponse(resp,false )
		//fmt.Println(string(response),i)

		//bytes, e := ioutil.ReadAll(resp.Body)

		//自动检查 编码类型
		//
		bodyreader :=bufio.NewReader(resp.Body)

		charset1:=determineEncoding(bodyreader)

		reader := transform.NewReader(bodyreader, charset1.NewDecoder())

		bytes, e := ioutil.ReadAll(reader)




		if e!=nil{
			panic(e)
		}


		//ParseCityList(bytes)
		fmt.Printf("%s",bytes)
		//fmt.Println(string(bytes),e)
	}else{

		fmt.Println("error  code = ",resp.StatusCode)
	}


}

// 根据html的meta头，试图自动转换编码到utf8
func determineEncoding(r *bufio.Reader) encoding.Encoding {
	data, err := r.Peek(1024)

	if err != nil {
		return unicode.UTF8
	}

	e, _, _ := charset.DetermineEncoding(data, "")
	return e
}

const cityListRegex = `<a href="(http://www.zhenai.com/zhenghun/\w+)"[^>]*>([^<]+)</a>`

func ParseCityList(contents []byte) {

	reg := regexp.MustCompile(cityListRegex)
	match := reg.FindAllSubmatch(contents, -1)
	for _, m := range match {
		for _,m2 :=range m{
			fmt.Print(string(m2))

		}
		fmt.Println()
	}

}
