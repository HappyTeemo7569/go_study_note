package main

import (
	"container/list"
	"fmt"
)

func main() {
	l := list.New()

	l.PushBack("test1")
	l.PushFront(11)
	showList(l)

	element := l.PushBack("test2")
	showList(l)
	l.InsertAfter("test_after_2", element)
	showList(l)
	l.InsertBefore("test_before_2", element)
	showList(l)
	l.Remove(element)
	showList2(l)

}

func showList(l *list.List) {
	fmt.Println("-------------------")
	e_tmp := l.Front()
	for i := 0; i < l.Len(); i++ {
		fmt.Println(e_tmp.Value)
		e_tmp = e_tmp.Next()
	}
	fmt.Println("-------------------")
}

func showList2(l *list.List) {
	fmt.Println("-------------------")
	for i := l.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
	fmt.Println("-------------------")
}
