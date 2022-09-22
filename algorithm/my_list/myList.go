package my_list

const (
	MAXSIZE = 20
)

type ElemType int //元素类型

const EmptyElement = 0 //空值

/**
线性表的抽象数据类型定义

定义：由零个（空表）或多个数据元素组成的序列

ADT:
    线性表List
Data:
    线性表的数据对象集合为{a1,a2,...,an},每个元素类型为DataType。除了第一个无前驱，最后一个无后继，
    其他每个元素都有一个字节前驱和直接后继结点。数据元素间关系一对一。
Operation:
    InitList(*L);//初始线性表,创建空表
    ClearList(*L);//清空线性表数据
    ListEmpty(L);//判断列表是否为空
    ListLength(L);//获取线性表的长度

    GetElem(L,i,* e);//获取指定位置的元素，返回在指针元素中
    LocateElem(L,e);//查找元素在线性表中的位置
    ListInsert(*L,i,e);//向线性表中指定位置插入元素
    ListDelete(*L, i, *e);//删除指定位置处的元素
*/

//注意线性表中的位置不是按照数组一样从0开始，而是按照我们正常习惯1开始的

type MyList interface {
	//四个基本操作，初始，清空，判断是否为空，获取长度
	InitList()
	ClearList()
	ListEmpty() bool
	ListLength() int

	//四个元素操作，插入，删除，两种查找
	GetElem(index int, e *ElemType) bool
	LocateElem(value ElemType) int
	ListInsert(index int, value ElemType) bool
	ListDelete(index int, e *ElemType) bool

	Echo()
	Test()
}
