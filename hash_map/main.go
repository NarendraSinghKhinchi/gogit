package main

type Node struct {
	key   int
	value int
	next  *Node
}

type MyHashMap struct {
	buckets    []*Node
	size       int
	capacity   int
	loadFactor float64
}

func Constructor() MyHashMap {
	return MyHashMap{
		buckets:    make([]*Node, 1000),
		size:       0,
		capacity:   1000,
		loadFactor: 0.7,
	}
}

func (this *MyHashMap) calcLoadFactor() float64 {
	return float64(this.size) / float64(this.capacity)
}

func (this *MyHashMap) hash(key int) int {
	return key % this.capacity
}

func (this *MyHashMap) rehash() {
	oldBuckets := this.buckets
	this.size = 0
	this.capacity *= 2
	this.buckets = make([]*Node, this.capacity)

	for _, node := range oldBuckets {
		for node != nil {
			this.Put(node.key, node.value)
			node = node.next
		}
	}

}

func (this *MyHashMap) Put(key int, value int) {
	index := this.hash(key)
	node := this.buckets[index]

	for node != nil {
		if node.key == key {
			node.value = value
			return
		}
		node = node.next
	}

	newNode := &Node{key: key, value: value}
	newNode.next = this.buckets[index]
	this.buckets[index] = newNode
	this.size++

	if this.calcLoadFactor() >= this.loadFactor {
		this.rehash()
	}
}

func (this *MyHashMap) Get(key int) int {
	index := this.hash(key)
	node := this.buckets[index]

	for node != nil {
		if node.key == key {
			return node.value
		}
		node = node.next
	}

	return -1
}

func (this *MyHashMap) Remove(key int) {
	index := this.hash(key)
	node := this.buckets[index]
	var nodePrev *Node

	for node != nil {
		if node.key == key {
			if nodePrev == nil {
				this.buckets[index] = node.next
			} else {
				nodePrev.next = node.next
			}
			this.size--
			return
		}
		nodePrev = node
		node = node.next
	}
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
