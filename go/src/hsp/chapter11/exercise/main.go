package main
import "fmt"
type Usb interface {
	Say()
}
type Stu struct {
	Age int
	Name string
	}
func (this *Stu) Say() {
	fmt.Println("Say()",this.Age)
	this.Age=12

}
func main() {
	var stu Stu = Stu{Age:10,Name:"wfh"}
	// 错误！ 会报 Stu类型没有实现Usb接口 , 
	// 如果希望通过编译,  var u Usb = &stu
	var u Usb = &stu
	//var u Usb = &stu
	u.Say()

	stu.Say()

	fmt.Println("here", u)
}

	