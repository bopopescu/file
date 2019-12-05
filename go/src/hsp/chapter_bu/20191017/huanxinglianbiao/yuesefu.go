package main

import (
	"fmt"
)

//小孩的结构体
type Boy struct {
	No   int  // 编号
	Next *Boy // 指向下一个小孩的指针[默认值是 nil]
}

// 编写一个函数，构成单向的环形链表
// num :表示小孩的个数
// *Boy : 返回该环形的链表的第一个小孩的指针
func AddBoy(num int) *Boy {
	first := &Boy{}  //空结点
	curBoy := &Boy{} //空结点
	//判断
	if num < 1 {
		fmt.Println("num 的值不对")
		return first
	}

	//循环的构建这个环形链表
	for i := 1; i <= num; i++ {
		boy := &Boy{No: i} //分析构成循环链表，需要一个辅助指针[帮忙的] //1. 因为第一个小孩比较特殊
		if i == 1 {        //第一个小孩
			first = boy //不要动
			curBoy = boy
			curBoy.Next = first //
		} else {
			curBoy.Next = boy
			curBoy = boy
			curBoy.Next = first //构造环形链表
		}
	}
	return first
}

//显示单向的环形链表[遍历]
func ShowBoy(first *Boy) {
	//处理一下如果环形链表为空
	if first.Next == nil {
		fmt.Println("链表为空，没有小孩...")
		return
	}
	//创建一个指针，帮助遍历.[说明至少有一个小孩]
	curBoy := first
	for {
		fmt.Printf("小孩编号=%d ->", curBoy.No) //退出的条件?curBoy.Next == first
		if curBoy.Next == first {
			break
		}
		//curBoy 移动到下一个
		curBoy = curBoy.Next
	}
}

/*
设编号为 1，2，... n 的 n 个人围坐一圈，约定编号为 k(1<=k<=n)
的人从 1 开始报数，数到 m 的那个人出列，它的下一位又从 1 开始报数，
数到 m 的那个人又出列，依次类推，直到所有人出列为止，由此产生一个出队编号的序列 */

func PlayGame(first *Boy, startNo int, countNum int) {
	//1. 空的链表我们单独的处理
	if first.Next == nil {
		fmt.Println("空的链表，没有小孩")
		return
	}
	//留一个，判断 startNO <= 小孩的总数
	//2. 需要定义辅助指针，帮助我们删除小孩
	tail := first
	//3. 让 tail 执行环形链表的最后一个小孩,这个非常的重要 //因为 tail 在删除小孩时需要使用到.
	for {
		if tail.Next == first { //说明 tail 到了最后的小孩
			break
		}
		tail = tail.Next
	}

	// now we can find the temp

	//first 可能会被干掉 所以不需要辅助节点了  直接用first节点不断改变


	//4. 让 first 移动到 startNo [后面我们删除小孩，就以 first 为准]
	for i := 1; i <= startNo-1; i++ {
		first = first.Next
		tail = tail.Next
	}
	fmt.Println()

//deprecated
	//5. 开始数 countNum, 然后就删除 first 指向的小孩
	for {


		//开始数 countNum-1 次
		for i := 1; i <= countNum-1; i++ {
			first = first.Next
			tail = tail.Next
		}
		fmt.Printf("小孩编号为%d 出圈 \n", first.No) //删除 first 执行的小孩
		first = first.Next
		tail.Next = first
		//判断如果 tail == first, 圈子中只有一个小孩.
		if tail == first {
			break
		}




		fmt.Println(first.Next == first)

	}
	fmt.Printf("小孩小孩编号为%d 出圈 \n", first.No)
		fmt.Println(first.Next == first)
}

func main() {






	first := AddBoy(500) //显示
	ShowBoy(first)
	PlayGame(first, 1, 1 )




//	first := &Boy{}    //空结点
//	curBoy := &Boy{}   //空结点
//
//
//
//	boy := &Boy{No: 1} //分析构成循环链表，需要一个辅助指针[帮忙的] //1. 因为第一个小孩比较特殊
//	fmt.Printf("%+v\n", boy)
//	return
//	first = boy        //不要动
//	curBoy = boy
//	fmt.Println(boy.Next)
//
//	curBoy.Next=first
//	fmt.Println(boy.Next)
//	fmt.Println(curBoy == first)
//
//	fmt.Println(first.No,first.Next,33)
//	fmt.Println(curBoy.No,curBoy,44)
//
//
//
//	boy2 := &Boy{No: 2} //分析构成循环链表，需要一个辅助指针[帮忙的] //1. 因为第一个小孩比较特殊
//	curBoy.No=3 //fmt.Println(curBoy == first)  true
//
//	curBoy= boy2  //fmt.Println(curBoy == first)  false
//fmt.Println(curBoy == first)
//
//	curBoy.Next = boy2
//	curBoy = boy2
//	curBoy.Next = first //构造环形链表
//	return
//
//
//	fmt.Println((&first))
//	fmt.Println((&curBoy))
//	fmt.Printf("%p", first)
//	fmt.Println()
//	fmt.Printf("%p", curBoy)
//
//
//	return


	//first := AddBoy(500) //显示
	//ShowBoy(first)
	//PlayGame(first, 20, 31)



	//return
	//first := &Boy{}    //空结点
	//curBoy := &Boy{}   //空结点
	//boy := &Boy{No: 1} //分析构成循环链表，需要一个辅助指针[帮忙的] //1. 因为第一个小孩比较特殊
	//first = boy        //不要动
	//curBoy = boy
	//
	//fmt.Println(first)
	//fmt.Printf("%p", first)
	//fmt.Println()
	//fmt.Println(curBoy)
	//fmt.Printf("%p", curBoy)
	//
	//fmt.Println()
	//fmt.Printf("%p", boy)

}
