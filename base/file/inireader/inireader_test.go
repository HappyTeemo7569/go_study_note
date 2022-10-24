package main

import (
	"fmt"
	"testing"
)

func Test_getValue(t *testing.T) {
	fmt.Println(getValue("./example.ini", "remote \"origin\"", "fetch"))
	fmt.Println(getValue("./example.ini", "core", "hideDotFiles"))
}
