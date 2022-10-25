package main //包名

import (
	_ "my_go/myLib"
	"my_go/web"
)

func webTest() {
	web.RunServer()
}
