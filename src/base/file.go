package base

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func CreateFile(path string) os.File {
	fout, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
	}
	return *fout
}

func WriteByte(f os.File, b []byte) {
	f.Write(b)
}

func WriteString(f os.File, input string) {
	f.WriteString(input)
}

func ReadFile(path string) {
	//打开文件
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}

	//关闭文件
	defer f.Close()

	buf := make([]byte, 1024*2) //2k大小

	//n代表从文件读取内容的长度
	n, err1 := f.Read(buf)
	if err1 != nil && err1 != io.EOF { //文件出错，同时没有到结尾
		fmt.Println("err1 = ", err1)
		return
	}
	fmt.Println("buf = ", string(buf[:n]))
}

//每次读取一行
func ReadFileLine(path string) {
	//打开文件
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}

	//关闭文件
	defer f.Close()

	//新建一个缓冲区，把内容先放在缓冲区
	r := bufio.NewReader(f)

	for {
		//遇到'\n'结束读取, 但是'\n'也读取进入
		buf, err := r.ReadBytes('\n')
		if err != nil {
			if err == io.EOF { //文件已经结束
				break
			}
			fmt.Println("err = ", err)
		}

		fmt.Printf("buf = #%s#\n", string(buf))
	}

}

func Test() {

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
