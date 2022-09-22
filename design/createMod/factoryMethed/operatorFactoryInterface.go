package factoryMethod

/**
工厂方法

工厂方法模式使用子类的方式延迟生成对象到子类中实现。
Go中不存在继承 所以使用匿名组合来实现
*/

//Operator 是被封装的实际类接口
type Operator interface {
	SetA(int)
	SetB(int)
	Result() int
}

//OperatorFactory 是工厂接口
type OperatorFactory interface {
	Create() Operator
}
