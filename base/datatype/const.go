package main

import "fmt"

func main() {
	//常量
	const a_con, b_con, c_con = 1, false, "str" //多重赋值
	println(a_con, b_con, c_con)

	// a_con = 3    //err, 常量不能修改

	//枚举
	const (
		Unknown = 0
		Female  = 1
		Male    = 2
	)

	/*
	   const (
	       a = iota   //0
	       b          //1
	       c          //2
	       d = "ha"   //独立值，iota += 1
	       e          //"ha"   iota += 1
	       f = 100    //iota +=1
	       g          //100  iota +=1
	       h = iota   //7,恢复计数
	       i          //8
	   )
	   fmt.Println(a,b,c,d,e,f,g,h,i)
	*/

	const (
		x = iota // x == 0
		y = iota // y == 1
		z = iota // z == 2
		w        // 这里隐式地说w = iota，因此w == 3。其实上面y和z可同样不用"= iota"
	)

	const v = iota // 每遇到一个const关键字，iota就会重置，此时v == 0
	const (
		h, i, j = iota, iota, iota //h=0,i=0,j=0 iota在同一行值相同
	)
	const (
		a       = iota //a=0
		b       = "B"
		c       = iota             //c=2
		d, e, f = iota, iota, iota //d=3,e=3,f=3 如果是同一行，值都一样
		g       = iota             //g = 4
	)
	const (
		x1 = iota * 10 // x1 == 0
		y1 = iota * 10 // y1 == 10
		z1 = iota * 10 // z1 == 20
	)

	fmt.Println(a, b, c, d, e, f, g, h, i) //0 1 2 ha ha 100 100 7 8
}
