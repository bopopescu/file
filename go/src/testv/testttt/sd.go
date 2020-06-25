package main

import (
	_"container/list"
	"fmt"
)

func main() {
	var ji int
	ji=1234

	var jikeofKejimeixue_of_daletou2000WsuperPrize = make(chan [3]int,20)

	var jit string = "asd"
	//var c1 byte = 'a'
	var c1 byte = '1'
	fmt.Println(jikeofKejimeixue_of_daletou2000WsuperPrize,ji,jit,c1)
	return
	//var masala fuck
	//
	//masala =fuck{"asd",bmw{123}}
	//
	//
	//fmt.Println(ji,masala)
	//
	//
	//var proche   List
	//proche = List{Element{},2}
	//
	//var jisd = List{Element{},23}
	//
	//
	//var jisdd = list.List{list.Element{},22}
	//dsd := list.List{list.Element{},23}
	//
	// jisdsds := List{Element{},23}
	//
	//var ui list.List
	//ui = list.List{list.Element{},33}
	//fmt.Println(proche,jisd,dsd,jisdd,ui,jisdsds)
}

type fuck struct{
	cc string
	benz  bmw

}

type bmw struct {
	ff int
}
type List struct {
	root Element // sentinel list element, only &root, root.prev, and root.next are used
	len  int     // current list length excluding (this) sentinel element
}
type Element struct {
	// Next and previous pointers in the doubly-linked list of elements.
	// To simplify the implementation, internally a list l is implemented
	// as a ring, such that &l.root is both the next element of the last
	// list element (l.Back()) and the previous element of the first list
	// element (l.Front()).
	next, prev *Element

	// The list to which this element belongs.
	list *List

	// The value stored with this element.
	Value interface{}
}
