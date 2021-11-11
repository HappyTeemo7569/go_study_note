package LRU

/**
新访问数据不断追加到 head 前边，旧数据不断从 tail 剔除。
LRU 使用链表顺序性保证了热数据在 head，冷数据在 tail。
*/

type Node struct {
	key        string // 淘汰 tail 时需在维护的哈希表中删除，不是冗余存储
	val        interface{}
	prev, next *Node // 双向指针
}

type List struct {
	head, tail *Node
	size       int // 缓存空间大小
}

//节点新增
func (l *List) Prepend(node *Node) *Node {
	if l.head == nil {
		l.head = node
		l.tail = node
	} else {
		node.prev = nil
		node.next = l.head
		l.head.prev = node
		l.head = node
	}
	l.size++
	return node
}

//节点移除
func (l *List) Remove(node *Node) *Node {
	if node == nil {
		return nil
	}
	prev, next := node.prev, node.next
	if prev == nil {
		l.head = next // 删除头结点
	} else {
		prev.next = next
	}

	if next == nil {
		l.tail = prev // 删除尾结点
	} else {
		next.prev = prev
	}

	l.size--
	node.prev, node.next = nil, nil
	return node
}

// 封装数据已存在缓存的后续操作
func (l *List) MoveToHead(node *Node) *Node {
	if node == nil {
		return nil
	}
	n := l.Remove(node)
	return l.Prepend(n)
}

func (l *List) Tail() *Node {
	return l.tail
}

func (l *List) Size() int {
	return l.size
}

/**
Set(k, v)
	*数据已缓存，则更新值，挪到 head 前
	*数据未缓存
		*缓存空间未满：直接挪到 head 前
		*缓存空间已满：移除 tail 并将新数据挪到 head 前

Get(k)
	*命中：节点挪到 head 前，并返回 value
	*未命中：返回 -1

*/

type LRUCache struct {
	capacity int // 缓存空间大小
	items    map[string]*Node
	list     *List
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		items:    make(map[string]*Node),
		list:     new(List),
	}
}

func (c *LRUCache) Set(k string, v interface{}) {
	// 命中
	if node, ok := c.items[k]; ok {
		node.val = v                         // 命中后更新值
		c.items[k] = c.list.MoveToHead(node) //
		return
	}

	// 未命中
	node := &Node{key: k, val: v} // 完整的 node
	if c.capacity == c.list.size {
		tail := c.list.Tail()
		delete(c.items, tail.key) // k-v 数据存储与 node 中
		c.list.Remove(tail)
	}
	c.items[k] = c.list.Prepend(node) // 更新地址
}

func (c *LRUCache) Get(k string) interface{} {
	node, ok := c.items[k]
	if ok {
		c.items[k] = c.list.MoveToHead(node)
		return node.val
	}
	return -1
}
