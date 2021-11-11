package algorithm

type Node struct {
	key        string // 淘汰 tail 时需在维护的哈希表中删除，不是冗余存储
	val        interface{}
	prev, next *Node // 双向指针
}

type List struct {
	head, tail *Node
	size       int // 缓存空间大小
}

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
