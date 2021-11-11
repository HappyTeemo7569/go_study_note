package LFU

type Node struct {
	key        string
	val        interface{}
	freq       int // 将节点从旧梯队移除时使用，非冗余存储
	prev, next *Node
}

type List struct {
	head, tail *Node
	size       int
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

type LFUCache struct {
	capacity int
	minFreq  int // 最低频率

	items map[string]*Node
	freqs map[int]*List // 不同频率梯队
}

/**
Set(k, v)
	*数据已缓存，则更新值，挪到下一梯队
	*数据未缓存
		*缓存空间未满：直接挪到第 1 梯队
		*缓存空间已满：移除 minFreq 梯队的 tail 节点，并将新数据挪到第 1 梯队


Get(k)
	*命中：节点挪到下一梯队，并返回 value
	*未命中：返回 -1
*/

func NewLFUCache(capacity int) *LFUCache {
	return &LFUCache{
		capacity: capacity,
		minFreq:  0,
		items:    make(map[string]*Node),
		freqs:    make(map[int]*List),
	}
}

func (c *LFUCache) Get(k string) interface{} {
	node, ok := c.items[k]
	if !ok {
		return -1
	}

	// 移到 +1 梯队中
	c.freqs[node.freq].Remove(node)
	node.freq++
	if _, ok := c.freqs[node.freq]; !ok {
		c.freqs[node.freq] = NewList()
	}
	newNode := c.freqs[node.freq].Prepend(node)
	c.items[k] = newNode // 新地址更新到 map
	if c.freqs[c.minFreq].Size() == 0 {
		c.minFreq++ // Get 的正好是当前值
	}
	return newNode.val
}

func (c *LFUCache) Set(k string, v interface{}) {
	if c.capacity <= 0 {
		return
	}

	// 命中，需要更新频率
	if val := c.Get(k); val != -1 {
		c.items[k].val = v // 直接更新值即可
		return
	}

	node := &Node{key: k, val: v, freq: 1}

	// 未命中
	// 缓存已满
	if c.capacity == len(c.items) {
		old := c.freqs[c.minFreq].Tail() // 最低最旧
		c.freqs[c.minFreq].Remove(old)
		delete(c.items, old.key)
	}

	// 缓存未满，放入第 1 梯队
	c.items[k] = node
	if _, ok := c.freqs[1]; !ok {
		c.freqs[1] = NewList()
	}
	c.freqs[1].Prepend(node)
	c.minFreq = 1
}
