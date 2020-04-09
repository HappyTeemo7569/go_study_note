package my_stack

const (
	MAXSIZE = 20
)

type ElemType int //元素类型

const EmptyElement = 0 //空值

/**
ADT 栈(stack)
Data
    同线性表。元素具有相同的类型，相邻元素具有前驱和后继的关系。

Operation
    InitStack( *s): 初始化操作，建立一个空栈
    ClearStack( *s): 将栈清空
    StackEmpty( s): 若栈存在，返回true，否则返回false
    StackLength( s): 返回栈S的元素个数

    GetTop( s, *e): 若是栈存在且非空，用e返回S的栈顶元素
    Push( *s, e):若是栈存在，则插入新的元素e到栈S中并成为栈顶元素
    Pop( *s, *e):若是栈存在且非空，删除栈顶元素，并用e返回其值
    DestroyStack( *s): 若是栈存在，则销毁他
endADT
*/

type MyTask interface {
	//四个基本操作，初始，清空，判断是否为空，获取长度
	InitStack()
	ClearStack()
	StackEmpty() bool
	StackLength() int

	//四个元素操作，插入，删除，两种查找
	GetTop(e *ElemType) bool
	Push(value ElemType) bool
	Pop(e *ElemType) bool
	DestroyStack() bool

	Echo()
	Test()
}
