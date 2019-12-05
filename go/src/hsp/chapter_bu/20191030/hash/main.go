package main

import (
	"fmt"
)

//定义 emp
type Emp struct {
	Id   int
	Name string
	Next *Emp
} //方法待定..

//定义 EmpLink
//我们这里的 EmpLink 不带表头,即第一个结点就存放雇员
type EmpLink struct {
	Head *Emp
}
//定义 hashtable ,含有一个链表数组
type HashTable struct {
	LinkArr [7]EmpLink
}
//方法待定..

//1. 添加员工的方法, 保证添加时，编号从小到大
func (this *EmpLink) Insert(emp *Emp) {
	cur := this.Head   // 这是辅助指针
	var pre *Emp = nil // 这是一个辅助指针 pre 在 cur 前面 //如果当前的 EmpLink 就是一个空链表
	if cur == nil {
		this.Head = emp //完成
		return
	}
	//如果不是一个空链表,给 emp 找到对应的位置并插入 //思路是 让 cur 和 emp 比较，然后让 pre 保持在 cur 前面
	for {
		if cur != nil { //比较
			if cur.Id > emp.Id { //找到位置
				break
			}
			pre = cur //保证同步
			cur = cur.Next
		} else {
			break
		}
	}
	//退出时，我们看下是否将 emp 添加到链表最后
	pre.Next = emp
	emp.Next = cur
}

//显示链表的信息
func (this *EmpLink) ShowLink(no int) {
	if this.Head == nil {
		fmt.Printf("链表%d 为空\n", no)

		return
	}
	//变量当前的链表，并显示数据
	cur := this.Head // 辅助的指针
	for {
		if cur != nil {
			fmt.Printf("链表%d 雇员 id=%d 名字=%s ->", no, cur.Id, cur.Name)
			cur = cur.Next
		} else {
			break
		}
	}
	fmt.Println() //换行处理
}



//给 HashTable 编写 Insert 雇员的方法.
func (this *HashTable) Insert(emp *Emp) {
	//使用散列函数，确定将该雇员添加到哪个链表
	linkNo := this.HashFun(emp.Id)   //使用对应的链表添加
	this.LinkArr[linkNo].Insert(emp) //
}

//编写方法，显示 hashtable 的所有雇员
func (this *HashTable) ShowAll() {
	for i := 0; i < len(this.LinkArr); i++ {
		this.LinkArr[i].ShowLink(i)
	}
}

func (this *HashTable) HashFun(id int) int {
	return id % 7 //得到一个值，就是对于的链表的下标
}

func (this *EmpLink) FindByid(id int) *Emp{

	cur := this.Head

	for {
		if cur!=nil && cur.Id==id{
			return cur
		}else if cur==nil{
			break
		}
		cur = cur.Next
	}
	return nil
}

func (this *HashTable) findByid(id int) *Emp{

	linkno :=this.HashFun(id)
	return this.LinkArr[linkno].FindByid(id)
}
func (this *Emp) showMe(){
	fmt.Println(this.Id,"is now found")
}
func main() {
	key := ""
	id := 0
	name := ""
	var hashtable HashTable
	for {
		fmt.Println("===============雇员系统菜单============")
		fmt.Println("input 表示添加雇员")
		fmt.Println("show 表示显示雇员")
		fmt.Println("find 表示查找雇员")
		fmt.Println("exit 表示退出系统")
		fmt.Println("请输入你的选择")
		fmt.Scanln(&key)
		switch key {
		case "input":
			fmt.Println("输入雇员 id")
			fmt.Scanln(&id)
			fmt.Println("输入雇员 name")
			fmt.Scanln(&name)
			emp := &Emp{

				Id:   id,
				Name: name}
			hashtable.Insert(emp)
		case "show":
			hashtable.ShowAll()
		case "find":
			fmt.Scanln(&id)
			emp := hashtable.findByid(id)
			if emp!=nil{
				emp.showMe()
			}else{
				fmt.Println("can not find id emp in hashtable")
			}
			//
		case "exit":
		default:
			fmt.Println("输入错误")
		}
	}
}
