package main
import (
	"fmt"
)



type Stu struct {
	Name string
}

func (stu Stu) Say() {
	fmt.Println("Stu Say()")
}


type integer int

func (i integer) Say() {
	fmt.Println("integer Say i =" ,i )
}


type AInterface interface {
	Say()
}

type BInterface interface {
	Hello()
}
type Monster struct {

	Name string
}
func (m Monster) Hello() {
	fmt.Println("Monster Hello()~~")
}

func (m Monster) Say() {
	fmt.Println("Monster Say()~~")
}



//内存对齐     golang-sizeof.tips
type A struct {
	X bool
	Y float64
	Z int16
	S string
	S2 string
}


func main() {
	var slice3 []A =make([]A,5)
	fmt.Println(slice3)
	fmt.Printf("%p\n",slice3)
	fmt.Printf("%p\n",&slice3[1])
	fmt.Printf("%p\n",&slice3[2])
	fmt.Printf("%p\n",&slice3[3])
	fmt.Printf("%p\n",&slice3[4])


	var sbool []bool =make([]bool,5)
	fmt.Println(sbool)
	fmt.Printf("%p\n",sbool)
	fmt.Printf("%p\n",&sbool[1])
	fmt.Printf("%p\n",&sbool[2])
	fmt.Printf("%p\n",&sbool[3])
	fmt.Printf("%p\n",&sbool[4])

	var slice []Monster =make([]Monster,5)
	slice[0] = Monster{Name:"asd"}
	slice[1] = Monster{Name:"asdd"}
	slice[4] = Monster{Name:"sdasasdd"}

	fmt.Println(slice)

	fmt.Printf("%p\n",slice)
	fmt.Printf("%p\n",&slice[1])
	fmt.Printf("%p\n",&slice[2])


	var slicetest = make([]uint32,5)
	slicetest[0] =123
	slicetest[1] = 255
	slicetest[4] = 4294967295


	fmt.Println(slicetest)

	fmt.Printf("%p\n",slicetest)
	fmt.Printf("%p\n",&slicetest[1])
	fmt.Printf("%p\n",&slicetest[2])
	fmt.Printf("%p\n",&slicetest[3])
	fmt.Printf("%p\n",&slicetest[4])


	var slicetest2 = make([]string,5)
	slicetest2[0] ="a"
	slicetest2[1] = "1"
	slicetest2[4] = "4294967295"


	fmt.Println(slicetest2)

	fmt.Printf("%p\n",slicetest2)
	fmt.Printf("%p\n",&slicetest2[1])
	fmt.Printf("%p\n",&slicetest2[2])
	return


	var stu Stu //结构体变量，实现了 Say() 实现了 AInterface
 	var a AInterface = stu
	a.Say()


	var i integer = 10
	var b AInterface = i
	b.Say() // integer Say i = 10


	//Monster实现了AInterface 和 BInterface
	var monster Monster
	monster.Name = "12"
	var a2 AInterface = monster
	var b2 BInterface = monster
	fmt.Println(a2,b2,monster)
	a2.Say()
	b2.Hello()
}