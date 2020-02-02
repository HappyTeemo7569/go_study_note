package sort

import "fmt"

func MySwap(x, y int) (int, int) {
	var c = x
	x = y
	y = c
	return x, y
}

func EchoArray(nums []int) {
	for _, num := range nums {
		fmt.Println(num)
	}
}
