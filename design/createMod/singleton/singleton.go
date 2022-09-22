package singleton

import "sync"

/**
单例模式
*/

//Singleton 懒汉式单例
type Singleton struct{}

var singleton *Singleton
var once sync.Once

//GetInstance 用于获取单例模式对象
func GetInstance() *Singleton {
	once.Do(func() {
		singleton = &Singleton{}
	})
	return singleton
}

// Singleton2 饿汉式单例
type Singleton2 struct{}

var singleton2 *Singleton2

func init() {
	singleton2 = &Singleton2{}
}

// GetInstance 获取实例
func GetInstance2() *Singleton2 {
	return singleton2
}
