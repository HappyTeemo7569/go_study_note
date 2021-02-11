package main

import (
	"bytes"
	"fmt"
	"io"
	"sync"
	"time"
)

/**
条件变量 sync.Cond 的主要作用并不是保证在同一时刻仅有一个线程访问某一个共享资源，而是在对应的共享资源状态发生变化时，通知其它因此而阻塞的线程。
*/

/**
type Cond struct {
  noCopy noCopy

  // L is held while observing or changing the condition
  L Locker

  notify  notifyList
  checker copyChecker
}

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
					fmt.Printf("reader success -%d: %s\n", i, data)
				}
				fmt.Println("null wait")
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

//单个
func main1() {
	db := NewDataBucket()
	go db.Read(1) // 开启读取器协程
	go func(i int) {
		d := fmt.Sprintf("data-%d", i)
		_, err := db.Put([]byte(d)) // 写入数据到缓冲区
		fmt.Println("write success")
		if err != nil {
			println(err.Error())
		}
	}(1) // 开启写入器协程
	time.Sleep(100 * time.Millisecond)
}

//多个
func main() {
	db := NewDataBucket()
	for i := 1; i < 3; i++ { // 启动多个读取器
		go db.Read(i)
	}
	for j := 0; j < 10; j++ { // 启动多个写入器
		go func(i int) {
			d := fmt.Sprintf("data-%d", i)
			_, err := db.Put([]byte(d)) // 写入数据到缓冲区
			fmt.Println("write success")
			if err != nil {
				println(err.Error())
			}
		}(j)
		time.Sleep(100 * time.Millisecond) // 每次启动一个写入器暂停100ms，让读取器阻塞
	}
}

/**
可以看到，通过互斥锁+条件变量，我们可以非常方便的实现多个 Go 协程之间的通信.
但是这个还是比不上 channel，因为 channel 还可以实现数据传递，条件变量只是发送信号，唤醒被阻塞的协程继续执行.
另外 channel 还有超时机制，不会出现协程等不到信号一直阻塞造成内存堆积问题，换句话说，channel 可以让程序更可控。
*/
