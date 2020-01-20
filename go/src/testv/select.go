package testv

import (
	"fmt"
	"strconv"
	"time"
)
func GetTime2( logTime, timeType string ) string {
	var item string
	switch timeType {
	case "day":
		item = "2006-01-02"
		break
	case "hour":
		item = "2006-01-02 15"
		break
	case "min":
		item = "2006-01-02 15:04"
		break
	}
	t, _ := time.Parse( item, "2014-06-15 08:37" )
	fmt.Println(t)
	return strconv.FormatInt( t.Unix(), 10 )
}

func main5() {

	Abc()
	fmt.Println(GetTime2("","day"))
	fmt.Println(GetTime2("","min"))

return

	ko := make(chan int ,5)

	for{

	select {

		case i,ok:=<-ko:

			if ok{
				fmt.Println(i)
			}else{
				fmt.Println("none")
			}
		default:
			fmt.Println(2222)
		}
	}


	time.Sleep(3e9)
}
