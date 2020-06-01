package goroutine

import (
	"bytes"
	"fmt"
	"io"
	"sync"
	"time"
)

/**
type Cond struct {
    noCopy noCopy

    // L is held while observing or changing the condition
    L Locker

    notify  notifyList
    checker copyChecker
}
*/

/**
// 等待通知
func (c *Cond) Wait() {
    c.checker.check()
    t := runtime_notifyListAdd(&c.notify)
    c.L.Unlock()
    runtime_notifyListWait(&c.notify, t)
    c.L.Lock()
}

// 单发通知
func (c *Cond) Signal() {
    c.checker.check()
    runtime_notifyListNotifyOne(&c.notify)
}

// 广播通知
func (c *Cond) Broadcast() {
    c.checker.check()
    runtime_notifyListNotifyAll(&c.notify)
}
*/

/**
声明一个Buffer的四种方法：
var b bytes.Buffer       //直接定义一个Buffer变量，不用初始化，可以直接使用
b := new(bytes.Buffer)   //使用New返回Buffer变量
b := bytes.NewBuffer(s []byte)   //从一个[]byte切片，构造一个Buffer
b := bytes.NewBufferString(s string)   //从一个string变量，构造一个Buffer

往Buffer中写入数据
b.Write(d []byte)        //将切片d写入Buffer尾部
b.WriteString(s string)  //将字符串s写入Buffer尾部
b.WriteByte(c byte)     //将字符c写入Buffer尾部
b.WriteRune(r rune)     //将一个rune类型的数据放到缓冲器的尾部
b.WriteTo(w io.Writer)  //将Buffer中的内容输出到实现了io.Writer接口的可写入对象中

注：将文件中的内容写入Buffer,则使用ReadForm(i io.Reader)


从Buffer中读取数据到指定容器
c := make([]byte,8)
b.Read(c)      //一次读取8个byte到c容器中，每次读取新的8个byte覆盖c中原来的内容
b.ReadByte()   //读取第一个byte，b的第1个byte被拿掉，赋值给a => a, _ := b.ReadByte()

b.ReadRune()   //读取第一个rune，b的第1个rune被拿掉，赋值给r => r, _ := b.ReadRune()
b.ReadBytes(delimiter byte)   //需要一个byte作为分隔符，读的时候从缓冲器里找第一个出现的分隔符（delim），找到后，把从缓冲器头部开始到分隔符之间的所有byte进行返回，作为byte类型的slice，返回后，缓冲器也会空掉一部分
b.ReadString(delimiter byte)  //需要一个byte作为分隔符，读的时候从缓冲器里找第一个出现的分隔符（delim），找到后，把从缓冲器头部开始到分隔符之间的所有byte进行返回，作为字符串返回，返回后，缓冲器也会空掉一部分
b.ReadForm(i io.Reader)  //从一个实现io.Reader接口的r，把r里的内容读到缓冲器里，n 返回读的数量
    file, _ := os.Open("./text.txt")
    buf := bytes.NewBufferString("Hello world")
    buf.ReadFrom(file)              //将text.txt内容追加到缓冲器的尾部
    fmt.Println(buf.String())

清空数据
b.Reset()

字符串化
b.String()
*/

// 数据 bucket
type DataBucket struct {
	buffer *bytes.Buffer //缓冲区
	mutex  *sync.RWMutex //互斥锁
	cond   *sync.Cond    //条件变量
}

func NewDataBucket() *DataBucket {
	buf := make([]byte, 0)
	db := &DataBucket{
		buffer: bytes.NewBuffer(buf),
		mutex:  new(sync.RWMutex),
	}
	db.cond = sync.NewCond(db.mutex.RLocker())
	return db
}

// 读取器
func (db *DataBucket) Read(i int) {
	db.mutex.RLock()         // 打开读锁
	defer db.mutex.RUnlock() // 结束后释放读锁
	var data []byte
	var d byte
	var err error
	for {
		//每次读取一个字节
		if d, err = db.buffer.ReadByte(); err != nil {
			if err == io.EOF { // 缓冲区数据为空时执行
				if string(data) != "" { // data 不为空，则打印它
					fmt.Printf("reader-%d: %s\n", i, data)
				}
				db.cond.Wait()  // 缓冲区为空，通过 Wait 方法等待通知，进入阻塞状态
				data = data[:0] // 将 data 清空
				continue
			}
		}
		data = append(data, d) // 将读取到的数据添加到 data 中
	}
}

// 写入器
func (db *DataBucket) Put(d []byte) (int, error) {
	db.mutex.Lock()         // 打开写锁
	defer db.mutex.Unlock() // 结束后释放写锁
	//写入一个数据块
	n, err := db.buffer.Write(d)
	db.cond.Signal() // 写入数据后通过 Signal 通知处于阻塞状态的读取器
	return n, err
}

func RunCond() {
	db := NewDataBucket()
	go db.Read(1) // 开启读取器协程
	go func(i int) {
		d := fmt.Sprintf("data-%d", i)
		db.Put([]byte(d)) // 写入数据到缓冲区
	}(1) // 开启写入器协程
	time.Sleep(100 * time.Millisecond)
}
