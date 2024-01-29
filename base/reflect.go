package main

import (
	"fmt"
	"reflect"
)

// 获取类型
func reflectTypeOf() {
	var a int
	//reflect.TypeOf() 取得变量 a 的类型对象 typeOfA，类型为 reflect.Type()。
	typeOfA := reflect.TypeOf(a)

	fmt.Println(typeOfA.Name(), typeOfA.Kind()) //int int

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
	//Map、Slice、Chan 属于引用类型，使用起来类似于指针，但是在种类常量定义中仍然属于独立的种类，不属于 Ptr。
	//type A struct{} 定义的结构体属于 Struct 种类，*A 属于 Ptr。

	type cat struct {
	}

	typeOfCat := reflect.TypeOf(cat{})
	fmt.Println(typeOfCat.Name(), typeOfCat.Kind()) //cat struct

	ins := &cat{}
	typeOfCatPtr := reflect.TypeOf(ins)
	fmt.Printf("name:'%v' kind:'%v'\n", typeOfCatPtr.Name(), typeOfCatPtr.Kind()) //name:'' kind:'ptr'
	//指针变量的类型名称是空，不是 *cat。

	// 取类型的元素
	typeOfCat2 := typeOfCatPtr.Elem() //等于取了指针的值 等效于	typeOfCat := reflect.TypeOf(cat{})
	// 显示反射类型对象的名称和种类
	fmt.Printf("element name: '%v', element kind: '%v'\n", typeOfCat2.Name(), typeOfCat2.Kind())

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
		//根据索引获取结构体字段类型
		/**
		type StructField struct {
			Name      string    // field name
			PkgPath   string    // PkgPath is the package path that qualifies a lower case (unexported)
								// 字段在结构体的路径
			Type      Type      // field type
			Tag       StructTag // field tag string
			Offset    uintptr   // offset within struct, in bytes
			Index     []int     // index sequence for Type.FieldByIndex
			Anonymous bool      // is an embedded field  是否是匿名
		}
		*/
		fieldType := typeOfCat.Field(i)
		// 输出成员名和tag
		fmt.Printf("name: %v  tag: '%v'\n", fieldType.Name, fieldType.Tag)
	}
	// 通过字段名, 找到字段类型信息
	if catType, ok := typeOfCat.FieldByName("Type"); ok {
		// 从tag中取出需要的tag
		fmt.Println(catType.Tag.Get("json"), catType.Tag.Get("id"))
	}

	valueOfCat := reflect.ValueOf(ins)
	//通过索引获取1这个位置的类型
	//不建议使用，一旦1这个位置没有，或者类型不对，就会panic
	fmt.Println(valueOfCat.FieldByIndex([]int{1}).Type())
}

// 重新获得原始值
func reflectValue() {
	var a int = 1024

	//获取变量的反射值对象
	valueOfA := reflect.ValueOf(a)

	//获取 interface { ｝类型的值，通过类型断 转换
	var getA int = valueOfA.Interface().(int)
	//获取 64 位的值，强制类型转换为 int 类型
	var getA2 int = int(valueOfA.Int())

	fmt.Println(getA, getA2)

}

//通过反射设置值
/**
1. 要可被寻址
*/
func reflectSetValue() {
	var a int = 1024

	//获取变量的反射值对象
	//valueOfA := reflect.ValueOf(a)
	//取出元素
	//valueOfAElem := valueOfA.Elem() //会报错，因为不可被寻址 Elem的参数需要是指针

	//获取变量的反射值对象 （a的地址）
	valueOfA := reflect.ValueOf(&a)
	//取出元素（a的值）
	valueOfAElem := valueOfA.Elem()

	fmt.Println(valueOfAElem.Int()) //1024
	valueOfAElem.SetInt(1)
	fmt.Println(valueOfAElem.Int())    //1
	fmt.Println(valueOfAElem.String()) //<int Value>
	//fmt.Println(valueOfAElem.Float()) // panic  call of reflect.Value.Float on int Value

}

func main() {
	reflectTypeOf()
	fmt.Println("-----------------")
	reflectStruct()
	fmt.Println("-----------------")
	reflectValue()
	fmt.Println("-----------------")
	reflectSetValue()
	fmt.Println("-----------------")
}
