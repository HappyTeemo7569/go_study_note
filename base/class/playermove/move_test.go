package main

import (
	"testing"
)

func Test_move(t *testing.T) {

	p := NewPlayer(0.5)
	brush := initBrush()

	p.MoveTo(Vec2{30, 10})
	p.Move(brush)

	p.MoveTo(Vec2{100, 100})
	p.Move(brush)

	p.MoveTo(Vec2{200, 50})
	p.Move(brush)

	p.MoveTo(Vec2{30, 200})
	p.Move(brush)

	brush.save()
}
