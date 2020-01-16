package base

import (
	"fmt"
)

type Animal struct {
	name   string
	weight int
}

type Horse struct {
	*Animal // 注意此行
	speak   string
}

func (h *Horse) hello() {
	fmt.Println(h.name)
	fmt.Println(h.weight)
	fmt.Println(h.speak)
}

func (a *Animal) hello() {
	fmt.Println(a.name)
	fmt.Println(a.weight)
	//fmt.Println(a.speak)
}

func TestClass() {
	bm_horse := &Horse{
		Animal: &Animal{ // 注意此行
			name:   "baima",
			weight: 60,
		},
		speak: "neigh",
	}
	bm_horse.hello()
	bm_horse.Animal.hello()
}
