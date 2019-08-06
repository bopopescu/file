package main

import "fmt"



func eval(a,b int ,op string) (int,error){
	switch op {
	case "+":
		return a+b,nil
	case "-":
		return a-b,nil
	default:
		return 0,fmt.Errorf(("no allowed op"+ op))
	}
}

func div(a,b int )(q,r int){
	return a/b,a%b
}


func main(){

	fmt.Println(eval(1,2,"1+"))

	var s,t = eval(333,222,"+2")
	fmt.Println(s,t)
	if t!=nil{
		fmt.Println("error:",t,s)
	}else{
		fmt.Println(s)
	}

	fmt.Println(int(23),float64(1222))
	var q,r = div(13,4)
	var first ,_= div(255,13)
	fmt.Println(q,r,first)
	fmt.Println("dsadas " + "dsadas")
}



