package main

import (
	"encoding/json"
	"fmt"
)

//type IT struct {
//	Company  string
//	Subjects []string
//	IsOk     bool
//	Price    float64
//}

type IT struct {
	//Company不会导出到JSON中
	Company string `json:"-"`

	// Subjects 的值会进行二次JSON编码
	Subjects []string `json:"subjects"`

	//转换为字符串，再输出
	IsOk bool `json:",string"`

	// 如果 Price 为空，则不输出到JSON串中
	Price float64 `json:"price, omitempty"`
}

func main() {

	t := IT{"itcast", []string{"Go", "C++", "Python", "Test"}, true, 666.666}

	//生成一段JSON格式的文本
	//如果编码成功， err 将赋于零值 nil，变量b 将会是一个进行JSON格式化之后的[]byte类型
	b, err := json.Marshal(t)
	//输出结果：{"Company":"itcast","Subjects":["Go","C++","Python","Test"],"IsOk":true,"Price":666.666}
	fmt.Println(string(b))

	b1, err := json.MarshalIndent(t, "", "    ")
	/*
	   输出结果：
	   {
	       "Company": "itcast",
	       "Subjects": [
	           "Go",
	           "C++",
	           "Python",
	           "Test"
	       ],
	       "IsOk": true,
	       "Price": 666.666
	   }
	*/
	if err != nil {
		fmt.Println("json err:", err)
	}
	fmt.Println(string(b1))

	//通过map生成JSON

	// 创建一个保存键值对的映射
	t1 := make(map[string]interface{})
	t1["company"] = "itcast"
	t1["subjects "] = []string{"Go", "C++", "Python", "Test"}
	t1["isok"] = true
	t1["price"] = 666.666

	b2, err := json.Marshal(t1)
	//json.MarshalIndent(t1, "", "    ")
	if err != nil {
		fmt.Println("json err:", err)
	}
	fmt.Println(string(b2))
	//输出结果：{"company":"itcast","isok":true,"price":666.666,"subjects ":["Go","C++","Python","Test"]}
}
