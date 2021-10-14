package main

import (
	"fmt"
	"time"
)

/**
type Time struct {
    //表示距离公元 1 年 1 月 1 日 00:00:00UTC 的秒数；
    wall uint64
    //表示纳秒；
    ext  int64
    //代表时区，主要处理偏移量，不同的时区，对应的时间不一样。
    loc *Location
}
*/

func main() {
	now := time.Now() //获取当前时间
	fmt.Printf("current time:%v\n", now)

	//获取年月日时分秒
	year := now.Year()     //年
	month := now.Month()   //月
	day := now.Day()       //日
	hour := now.Hour()     //小时
	minute := now.Minute() //分钟
	second := now.Second() //秒
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)

	//获取时间戳
	timestamp1 := now.Unix()     //时间戳
	timestamp2 := now.UnixNano() //纳秒时间戳
	fmt.Printf("现在的时间戳：%v\n", timestamp1)
	fmt.Printf("现在的纳秒时间戳：%v\n", timestamp2)

	//获取星期几
	fmt.Println(now.Weekday().String())

	//时间加减

	later := now.Add(time.Hour) // 当前时间加1小时后的时间
	fmt.Println(later)

	before := now.Before(later) // 当前时间和一小时后比较
	fmt.Println(before)

	after := now.After(later) // 当前时间和一小时后比较
	fmt.Println(after)

	//定时器
	ticker := time.Tick(time.Second) //定义一个1秒间隔的定时器
	for i := range ticker {
		fmt.Println(i) //每秒都会执行的任务
	}

	// 格式化的模板为Go的出生时间2006年1月2号15点04分 Mon Jan
	// 24小时制
	fmt.Println(now.Format("2006-01-02 15:04:05.000 Mon Jan"))
	// 12小时制
	fmt.Println(now.Format("2006-01-02 03:04:05.000 PM Mon Jan"))
	fmt.Println(now.Format("2006/01/02 15:04"))
	fmt.Println(now.Format("15:04 2006/01/02"))
	fmt.Println(now.Format("2006/01/02"))
}
