package my_array

import (
	"errors"
	"fmt"
)


type Array struct {
	data   []int
	newData []int
	length int
	free int
}

//为数组初始化内存
func NewArray(capacity int) *Array {
	if capacity == 0 {
		return nil
	}
	return &Array{
		data:   make([]int, capacity, capacity),
		length: 0,
		free:capacity,
	}
}

func (this *Array) Len() int {
	return this.length
}

func (this *Array) Free() int   {
	return this.free
}

//判断索引是否越界
func (this *Array) isIndexOutOfRange(index int) bool {
	if index >= this.Len() {
		return true
	}
	return false
}

//通过索引查找数组，索引范围[0,n-1]
func (this *Array) Find(index int) (int, error) {
	if this.isIndexOutOfRange(index) {
		return 0, errors.New("out of index range")
	}
	return this.data[index], nil
}

func (this *Array)kuoRong()  {
	this.newData = make([]int,2*this.length)
	this.free = this.length

	for k,v :=range this.data  {
		this.newData[k] = v
	}
	this.data = make([]int,2*this.length)
	this.data = this.newData
}

//插入数值到索引index上
func (this *Array) Insert(index int, v int) error {
	if this.free == 0 {
		this.kuoRong()
		//return errors.New("full array")
	}
	if index != this.length && this.isIndexOutOfRange(index) {
		return errors.New("out of index range")
	}

	for i := this.length; i > index; i-- {
		this.data[i] = this.data[i-1]
	}
	this.data[index] = v
	this.length++
	this.free--
	return nil
}

//插入到尾端
func (this *Array) InsertToTail(v int) error {
	return this.Insert(this.Len(), v)
}

//删除索引index上的值
func (this *Array) Delete(index int) (int, error) {
	if this.isIndexOutOfRange(index) {
		return 0, errors.New("out of index range")
	}
	v := this.data[index]
	for i := index; i < this.Len()-1; i++ {
		this.data[i] = this.data[i+1]
	}
	this.length--
	return v, nil
}

//打印数列
func (this *Array) Print() {
	var format string
	for i := 0; i < this.Len(); i++ {
		format += fmt.Sprintf("|%+v", this.data[i])
	}
	fmt.Println(format)
}
