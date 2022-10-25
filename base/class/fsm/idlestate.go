package main

import "fmt"

// 闲置状态
type IdleState struct {
	StateInfo // 使用StateInfo实现基础接口
}

// 重新实现状态开始
func (i *IdleState) OnBegin() {
	fmt.Println("IdleState begin")
}

// 重新实现状态结束
func (i *IdleState) OnEnd() {
	fmt.Println("IdleState end")
}
