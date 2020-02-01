package list

const (
	MAXSIZE = 20
)

type ElemType int

const EmptyElement = 0 //空值

type MyList interface {
	InitList()
	ClearList()
	CheckEmpty() bool
	GetLength() int
	AddItem(index int, value ElemType) bool
	DelItem(index int) bool
	SetItem(index int, value ElemType) bool
	GetItem(index int) ElemType
	Echo()
}
