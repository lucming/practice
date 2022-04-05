package leetcode

//LRU算法
type LRUCache struct {
	Cache      map[int]*Node
	Head, Tail *Node
	Size, Cap  int
}

type Node struct {
	Key, Value int
	Pre, Next  *Node
}

func NewNode(key, val int) *Node {
	return &Node{
		Key:   key,
		Value: val,
	}
}

func Constructor(capacity int) LRUCache {
	c := LRUCache{
		Cache: make(map[int]*Node),
		Head:  NewNode(0, 0),
		Tail:  NewNode(0, 0),
		Cap:   capacity,
	}
	c.Head.Next = c.Tail
	c.Tail.Pre = c.Head

	return c
}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.Cache[key]; !ok {
		return -1
	}

	node := this.Cache[key]
	this.moveToHead(node)

	return node.Value
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.Cache[key]; !ok {
		s := NewNode(key, value)
		this.addToHead(s)
		this.Cache[key] = s
		this.Size++
		if this.Size > this.Cap {
			moved := this.moveTail()
			delete(this.Cache, moved.Key)
			this.Size--
		}
	} else {
		node := this.Cache[key]
		node.Value = value
		this.moveToHead(node)
	}
}

func (this *LRUCache) moveToHead(p *Node) {
	this.moveNode(p)
	this.addToHead(p)
}

func (this *LRUCache) moveNode(p *Node) {
	p.Pre.Next = p.Next
	p.Next.Pre = p.Pre
}

func (this *LRUCache) addToHead(p *Node) {
	p.Pre = this.Head
	p.Next = this.Head.Next
	this.Head.Next.Pre = p
	this.Head.Next = p
}

func (this *LRUCache) moveTail() *Node {
	deleted := this.Tail.Pre
	this.moveNode(deleted)
	return deleted
}
