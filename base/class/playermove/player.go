package main

import "fmt"

type Player struct {
	currPos   Vec2    // 当前位置
	targetPos Vec2    // 目标位置
	speed     float32 // 移动速度
}

// 移动到某个点就是设置目标位置
func (p *Player) MoveTo(v Vec2) {
	p.targetPos = v
}

// 获取当前的位置
func (p *Player) Pos() Vec2 {
	return p.currPos
}

// 是否到达
func (p *Player) IsArrived() bool {
	// 通过计算当前玩家位置与目标位置的距离不超过移动的步长，判断已经到达目标点
	return p.currPos.DistanceTo(p.targetPos) < p.speed
}

// 逻辑更新
func (p *Player) Update() {

	if !p.IsArrived() {
		// 计算出当前位置指向目标的朝向
		dir := p.targetPos.Sub(p.currPos).Normalize()
		// 添加速度矢量生成新的位置
		newPos := p.currPos.Add(dir.Scale(p.speed))
		// 移动完成后，更新当前位置
		p.currPos = newPos
	}

}

//移动
func (p *Player) Move(brush *Brush) {
	for !p.IsArrived() {
		// 更新玩家位置
		p.Update()
		// 打印每次移动后的玩家位置
		fmt.Println(p.Pos())
		brush.draw(int(p.Pos().X), int(p.Pos().Y))
	}
}

// 创建新玩家
func NewPlayer(speed float32) *Player {
	return &Player{
		speed: speed,
	}
}
