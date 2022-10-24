package file

import (
	"fmt"
	"testing"
)

func Test_File(t *testing.T) {
	fout := CreateFile("./test2.txt")

	defer fout.Close() //main函数结束前， 关闭文件

	outstr := fmt.Sprintf("%s:%d\n", "Hello go", 1)
	WriteString(fout, outstr)         //写入string信息到文件
	WriteByte(fout, []byte("abcd\n")) //写入byte类型的信息到文件

	//将标准输入写入文件
	var a string
	fmt.Println("请输入a: ")
	fmt.Scan(&a) //从标准输入设备中读取内容，放在a中
	fmt.Println("a = ", a)
	WriteString(fout, a)

	//读文件
	ReadFile("./test.txt")
	ReadFileLine("./test.txt")
}
