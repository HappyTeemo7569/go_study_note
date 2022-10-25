package main

import (
	"fmt"
)

/**
有限状态机（ Finite State Machine, FSM ）， 表示有限个状态及在这些状态间的转移和动作等行为的数学模型。
本例将实现状态接口、状态管理器及一系列的状态和使用状态的逻辑。

状态的概念
状态机中的状态与状态问能够自由转换。但是现实当中的状态却不一定能够自由转换，
例如 人可以从站立状态转移到卧倒状态，却不能从卧倒状态直接转移到跑步状态，需要先经过站立状态后再转移到跑步状态。
每个状态可以设置它可以转移到的状态，有些状态机还允许在同个状态间互相转换，这也需要根据实际情况进行配置。
*/

// 封装转移状态和输出日志
func transitAndReport(sm *StateManager, target string) {
	if err := sm.Transit(target); err != nil {
		fmt.Printf("FAILED! %s --> %s, %s\n\n", sm.CurrState().Name(), target, err.Error())
	}
}

func main() {
	// 实例化一个状态管理器
	sm := NewStateManager()

	// 响应状态转移的通知
	sm.OnChange = func(from, to State) {
		// 打印状态转移的流向
		fmt.Printf("%s ---> %s\n\n", StateName(from), StateName(to))
	}

	// 添加3个状态
	sm.Add(new(IdleState))
	sm.Add(new(MoveState))
	sm.Add(new(JumpState))

	// 在不同状态间转移
	transitAndReport(sm, "IdleState")
	transitAndReport(sm, "MoveState")
	transitAndReport(sm, "MoveState")
	transitAndReport(sm, "JumpState")
	transitAndReport(sm, "JumpState")
	transitAndReport(sm, "IdleState")
}
