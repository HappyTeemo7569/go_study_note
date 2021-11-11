package LRU

import (
	"fmt"
	"testing"
)

func TestLRU(t *testing.T) {
	K1 := "string1"
	K2 := "string2"
	K3 := "string3"
	K4 := "string4"

	c := NewLRUCache(2)
	c.Set(K1, 1)
	c.Set(K2, 2)
	c.Set(K1, 100)
	fmt.Println(c.Get(K1)) // 100
	c.Set(K3, 3)
	fmt.Println(c.Get(K2)) // -1
	c.Set(K4, 4)
	fmt.Println(c.Get(K1)) // -1
	fmt.Println(c.Get(K3)) // 3
	fmt.Println(c.Get(K4)) // 4
}
