package main

import "fmt"

func main() {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(len(slice), cap(slice)) //10 10

	s1 := slice[2:5]
	fmt.Println(s1)               //[2 3 20]
	fmt.Println(len(s1), cap(s1)) //3 8

	s2 := s1[2:6:7]               // 4,5,6,7
	fmt.Println(len(s2), cap(s2)) // 4 5

	s2 = append(s2, 100)
	s2 = append(s2, 200)
	fmt.Println(s2) //[4 5 6 7 100 200]

	//0, 1, 2, 3(赋值), 4, 5, 6, 7, 8, 9

	s1[2] = 20

	fmt.Println(slice) //[0 1 2 3 20 5 6 7 100 9]

	delete_index := 2
	//查看删除位置之前的元素和 之后的元素
	fmt.Println(slice[:delete_index], slice[delete_index+1:])
	//将删除点前后的元素连接起来
	slice = append(slice[:delete_index], slice[delete_index+1:]...)
	fmt.Println(slice)
}
