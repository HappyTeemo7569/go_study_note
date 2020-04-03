package myLib

import (
	"fmt"
	"gopkg.in/gcfg.v1"
)

func InitConfig() {
	err := gcfg.ReadFileInto(&config, "conf/config.ini")
	if err != nil {
		Logger.Errorf("Failed to parse config file:", err)
		return
	}
	fmt.Println("调试模式:", config.Other.Debug)
}
