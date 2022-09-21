package main

import (
	"fmt"
	"reflect"
)

func main() {
	reflectStruct()
}
func reflectBase() {
	var a int
	//reflect.TypeOf() 取得变量 a 的类型对象 typeOfA，类型为 reflect.Type()。
	typeOfA := reflect.TypeOf(a)
	//类型名为 int，种类（Kind）为 int。
	fmt.Println(typeOfA.Name(), typeOfA.Kind())

	/**
	type Kind uint
	const (
	    Invalid Kind = iota  // 非法类型
	    Bool                 // 布尔型
	    Int                  // 有符号整型
	    Int8                 // 有符号8位整型
	    Int16                // 有符号16位整型
	    Int32                // 有符号32位整型
	    Int64                // 有符号64位整型
	    Uint                 // 无符号整型
	    Uint8                // 无符号8位整型
	    Uint16               // 无符号16位整型
	    Uint32               // 无符号32位整型
	    Uint64               // 无符号64位整型
	    Uintptr              // 指针
	    Float32              // 单精度浮点数
	    Float64              // 双精度浮点数
	    Complex64            // 64位复数类型
	    Complex128           // 128位复数类型
	    Array                // 数组
	    Chan                 // 通道
	    Func                 // 函数
	    Interface            // 接口
	    Map                  // 映射
	    Ptr                  // 指针
	    Slice                // 切片
	    String               // 字符串
	    Struct               // 结构体
	    UnsafePointer        // 底层指针
	)
	*/
	//Map、Slice、Chan 属于引用类型，使用起来类似于指针，但是在种类常量定义中仍然属于独立的种类，不属于 Ptr。type A struct{} 定义的结构体属于 Struct 种类，*A 属于 Ptr。

	type cat struct {
	}

	typeOfCat := reflect.TypeOf(cat{})
	fmt.Println(typeOfCat.Name(), typeOfCat.Kind()) //cat struct

	ins := &cat{}
	typeOfCatPtr := reflect.TypeOf(ins)
	fmt.Printf("name:'%v' kind:'%v'\n", typeOfCatPtr.Name(), typeOfCatPtr.Kind()) //指针变量的类型名称是空，不是 *cat。

	// 取类型的元素
	typeOfCat = typeOfCat.Elem() //等于取了指针的值 等效于	typeOfCat := reflect.TypeOf(cat{})
	// 显示反射类型对象的名称和种类
	fmt.Printf("element name: '%v', element kind: '%v'\n", typeOfCat.Name(), typeOfCat.Kind())

}

func reflectStruct() {

	type cat struct {
		Name string
		Type int `json:"type" id:"100"`
	}

	ins := cat{Name: "mimi", Type: 1}
	typeOfCat := reflect.TypeOf(ins)

	// 遍历结构体所有成员
	for i := 0; i < typeOfCat.NumField(); i++ {
		// 获取每个成员的结构体字段类型
		fieldType := typeOfCat.Field(i)
		// 输出成员名和tag
		fmt.Printf("name: %v  tag: '%v'\n", fieldType.Name, fieldType.Tag)
	}
	// 通过字段名, 找到字段类型信息
	if catType, ok := typeOfCat.FieldByName("Type"); ok {
		// 从tag中取出需要的tag
		fmt.Println(catType.Tag.Get("json"), catType.Tag.Get("id"))
	}
}
