package main

import (
	"fmt"
)

// 学生
type student struct {
	Name  string
	Age   int
	Score float64
}

func (s *student) showInfo() {
	fmt.Println("", s)
}

// 小学生
type pupil struct {
	student // 嵌入匿名结构体,就是继承了!
}

func (s *pupil) exam() {
	fmt.Println("小学生正在考试", s)
}

// 大学生
type univercityStudent struct {
	student // 嵌入匿名结构体,就是继承了!
}

func (s *univercityStudent) exam() {
	fmt.Println("大学生正在考试", s)
}

// 外星人
type Alia struct {
	stu student           // 嵌入有名结构体,就是继承了!
	uni univercityStudent // 多重继承
	int                   // 匿名字段也是可以的
}

func main() {

	da := &univercityStudent{}
	da.student.Age = 21
	// 也可以简化书写, 就近原则,它会先去子类查找,找不到了再去匿名结构体里查找.这就是覆盖把!
	da.Age = 25
	fmt.Println("", da)

	// 有名结构体(组合关系),此时不能简写了
	al := &Alia{}
	al.stu.Age = 1000
	fmt.Println("", al)
}
