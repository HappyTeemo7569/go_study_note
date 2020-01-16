package main //包名
import "base"

/**
fout := myLib2.CreateFile("./test2.txt")

	defer fout.Close() //main函数结束前， 关闭文件

	outstr := fmt.Sprintf("%s:%d\n", "Hello go", 1)
	myLib2.WriteString(fout, outstr)         //写入string信息到文件
	myLib2.WriteByte(fout, []byte("abcd\n")) //写入byte类型的信息到文件

	//将标准输入写入文件
	var a string
	fmt.Println("请输入a: ")
	fmt.Scan(&a) //从标准输入设备中读取内容，放在a中
	fmt.Println("a = ", a)
	myLib2.WriteString(fout, a)

	//读文件
	myLib2.ReadFile("./test.txt")
	myLib2.ReadFileLine("./test.txt")
*/
func main() {

	base.HelloWord()
	base.DataType()
}
