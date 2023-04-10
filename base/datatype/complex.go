package main

import "fmt"

/** 复数 **/

func main() {

	var comv1 complex64 // 由2个float32构成的复数类型
	comv1 = 3.2 + 12i
	comv2 := 3.2 + 12i        // v2是complex128类型
	comv3 := complex(3.2, 12) // v3结果同v2

	fmt.Println(comv1, comv2, comv3)
	//内置函数real(v1)获得该复数的实部
	//通过imag(v1)获得该复数的虚部
	fmt.Println(real(comv1), imag(comv1))
}