type LRUCache struct {
	Dummy    *Node
	Len, Cap int
	nodeMap  map[int]*Node
}

type Node struct {
	Key, Val   int
	Prev, Next *Node
}

func Constructor(capacity int) LRUCache {
	cache := LRUCache{
		Dummy:   new(Node),
		Cap:     capacity,
		nodeMap: make(map[int]*Node),
	}
	cache.Dummy.Prev, cache.Dummy.Next = cache.Dummy, cache.Dummy
	return cache
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.nodeMap[key]; ok {
		this.moveToHead(node)
		return node.Val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.nodeMap[key]; ok {
		node.Val = value
		this.moveToHead(node)
	} else {
		if this.Len == this.Cap {
			delete(this.nodeMap, this.Dummy.Prev.Key)
			this.deleteNode(this.Dummy.Prev)
		}
		n := &Node{Key: key, Val: value}
		this.insertHead(n)
		this.nodeMap[key] = n
	}
}

func (this *LRUCache) moveToHead(node *Node) {
	this.deleteNode(node)
	this.insertHead(node)
}

func (this *LRUCache) deleteNode(node *Node) {
	node.Prev.Next, node.Next.Prev = node.Next, node.Prev
	this.Len--
}

func (this *LRUCache) insertHead(node *Node) {
	node.Next, node.Prev = this.Dummy.Next, this.Dummy
	this.Dummy.Next, this.Dummy.Next.Prev = node, node
	this.Len++
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
