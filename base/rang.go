package main

import "fmt"

func main() {
	s := "abc"
	for i := range s { //支持 string/array/slice/map。
		fmt.Printf("%c\n", s[i])
	}

	for _, c := range s { // 忽略 index
		fmt.Printf("%c\n", c)
	}

	for i, c := range s {
		fmt.Printf("%d, %c\n", i, c)
	}
}
