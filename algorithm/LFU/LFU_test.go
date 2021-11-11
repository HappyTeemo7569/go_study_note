package LFU

import (
	"fmt"
	"testing"
)

func TestLFU(t *testing.T) {

	K1 := "string1"
	K2 := "string2"
	K3 := "string3"
	K4 := "string4"

	c := NewLFUCache(2)
	c.Set(K1, 1)           // 1:K1
	c.Set(K2, 2)           // 1:K2->K1
	fmt.Println(c.Get(K1)) // 1:K2 2:K1 // 1
	c.Set(K3, 3)           // 1:K3 2:K1
	fmt.Println(c.Get(K2)) // -1
	fmt.Println(c.Get(K3)) // 2:k3->k1 // 3
	c.Set(K4, 4)           // 1:K4 2:K3
	fmt.Println(c.Get(K1)) // -1
	fmt.Println(c.Get(K3)) // 1:K4 3:K3 // 3
}
