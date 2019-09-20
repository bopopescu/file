package main

import (
	"fmt"
	"strconv"
	"strings"
)

func buildin(){
	var str string = "hello我"
	// 含有中文的话要转成切片[]rune
	v := []rune(str)
	fmt.Println("字符串的长度是",len(v))

	// new用来给基本类型分配地址,make用来给引用类型分配地址
	num := new(int)
	fmt.Printf("num的类型是%T num的值是%v num的地址是%v num这个指针指向的值是%v",
		num,num,&num,*num)


}
func str2int(){
	n,err := strconv.Atoi("123q")
	if err != nil {
		fmt.Println("转换出错",err)
	}else {
		fmt.Println("结果是",n)
	}
}
func int2str(){
	n := strconv.Itoa(123)
	fmt.Println("结果是",n)
}
func main(){
	// 1,内嵌函数buildin,不用导包可以直接使用的
	buildin()
	// 2,字符串转整形
	//str2int()
	// 3,整形字符串
	//int2str()
	// 4,字符串转byte切片
	// var byt = []byte("hello")
	// fmt.Printf("%v",byt)

	// 5,10进制转2,8,16
	// res := strconv.FormatInt(33,2)
	// fmt.Println("33对应的二进制是",res)

	// 6,查找字符串中是否含有某个子串
	// b := strings.Contains("iloveyou","you")
	// fmt.Printf("%v",b)

	// 7,统计一个字符串中含有多少个子串
	// count := strings.Count("iloveyou","o")
	// fmt.Println("含有多少个o",count)

	// 8,比较字符串是否相等,EqualFold不区分大小写的. ==是区分大小写的.
	// b1 := strings.EqualFold("abc","Abc")
	// fmt.Printf("%v",b1)

	// 9,返回子串在字符串中第一次出现的index值
	// index := strings.Index("iloveyouabc","abc")
	// fmt.Printf("%v",index)

	// 10,返回子串在字符串中最后出现的index值
	// index1 := strings.LastIndex("iloveyouabffffabc","abc")
	// fmt.Printf("%v",index1)

	// 11,替换字符串
	// str := strings.Replace("iloveyou","you","me",-1)
	// fmt.Println("",str)

	// 12,拆分
	// arr := strings.Split("hello,work,ok",",")
	// fmt.Printf("%v",arr)

	// 13,大小写转换,ToUpper
	// st := strings.ToLower("HELlo")
	// fmt.Println("",st)

	// 14,去掉字符串左右的空格
	// s := strings.TrimSpace("  I love he  sister ")
	// fmt.Printf("str=%v",s)

	// 15,将左右两边指定的字符去掉
	// t := strings.Trim("!he!llo!","!")
	// fmt.Printf("str=%v",t)

	// 15,将左边指定的字符去掉,同理也可以去掉右边的
	s1 := strings.TrimLeft("!he!llo!","!")
	fmt.Printf("str=%v",s1)

	// 16,判断字符串是否以指定的字符开始或者结尾.HasSuffix
	b1 := strings.HasPrefix("hello","h")
	fmt.Printf("str=%v",b1)

}