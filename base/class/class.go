package class

import (
	"fmt"
)

/**
Go实现类似继承的例子
*/

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
}

func TestClass() {
	bmHorse := &Horse{
		Animal: &Animal{ // 注意此行
			name:   "baima",
			weight: 60,
		},
		speak: "neigh",
	}
	bmHorse.hello()
	bmHorse.Animal.hello()
}
