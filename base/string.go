package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	//判断是否包含
	res := strings.Contains("abcdefg", "abc")
	fmt.Println(res) //true

	//查找索引
	fmt.Println(strings.Index("chicken", "ken")) //4

	//字符串重复
	fmt.Println("ba" + strings.Repeat("na", 2)) //banana

	//字符串替换
	fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2)) //oinky oinky oink
	fmt.Println(strings.Replace("oink oink oink", "oink", "moo", -1))

	//数组和字符串互转
	s := []string{"foo", "bar", "baz"}
	fmt.Println(strings.Join(s, ", "))              //foo, bar, baz
	fmt.Printf("%q\n", strings.Split("a,b,c", ",")) //["a" "b" "c"]

	//字符串清理
	fmt.Printf("[%q]", strings.Trim(" !!! Achtung !!! ", "! ")) //["Achtung"]

	//去除空格，返回slice
	fmt.Printf("Fields are: %q", strings.Fields("  foo bar  baz   ")) //]Fields are: ["foo" "bar" "baz"]

	//Append 系列函数将整数等转换为字符串后，添加到现有的字节数组中。
	str := make([]byte, 0, 100)
	str = strconv.AppendInt(str, 4567, 10) //以10进制方式追加
	str = strconv.AppendBool(str, false)
	str = strconv.AppendQuote(str, "abcdefg")
	str = strconv.AppendQuoteRune(str, '单')
	fmt.Println(string(str)) //4567false"abcdefg"'单'

	//Format 系列函数把其他类型的转换为字符串。
	a := strconv.FormatBool(false)
	b := strconv.FormatInt(1234, 10)
	c := strconv.FormatUint(12345, 10)
	d := strconv.Itoa(1023)
	fmt.Println(a, b, c, d) //false 1234 12345 1023

	//Parse 系列函数把字符串转换为其他类型。
	a1, err := strconv.ParseBool("false")
	checkError(err)
	b1, err := strconv.ParseFloat("123.23", 64)
	checkError(err)
	c1, err := strconv.ParseInt("1234", 10, 64)
	checkError(err)
	d1, err := strconv.ParseUint("12345", 10, 64)
	checkError(err)
	e1, err := strconv.Atoi("1023")
	checkError(err)
	fmt.Println(a1, b1, c1, d1, e1) //false 123.23 1234 12345 1023

}

func checkError(e error) {
	if e != nil {
		fmt.Println(e)
	}
}
