package lru

import "fmt"

type LruCache interface {
	Put(key, value string)
	Get(key string) (string, bool)
}

type node struct {
	key   string
	value string
	prev  *node
	next  *node
}

type lruCache struct {
	capacity int
	items    map[string]*node

	head *node
	tail *node
}

func NewLruCache(capacity int) LruCache {
	if capacity <= 0 {
		panic(fmt.Sprintf("lru cache capacity must be > 0, got %d", capacity))
	}

	return &lruCache{
		capacity: capacity,
		items:    make(map[string]*node, capacity),
	}
}

func (c *lruCache) Put(key, value string) {
	if n, ok := c.items[key]; ok {
		n.value = value
		c.moveToFront(n)
		return
	}

	if len(c.items) >= c.capacity {
		c.evictLeastRecentlyUsed()
	}

	n := &node{key: key, value: value}
	c.items[key] = n
	c.addToFront(n)
}

func (c *lruCache) Get(key string) (string, bool) {
	n, ok := c.items[key]
	if !ok {
		return "", false
	}

	c.moveToFront(n)
	return n.value, true
}

func (c *lruCache) evictLeastRecentlyUsed() {
	if c.tail == nil {
		return
	}

	lru := c.tail
	c.removeNode(lru)
	delete(c.items, lru.key)
}

func (c *lruCache) addToFront(n *node) {
	if c.head == nil {
		c.head = n
		c.tail = n
		return
	}

	n.next = c.head
	c.head.prev = n
	c.head = n
}

func (c *lruCache) moveToFront(n *node) {
	if n == c.head {
		return
	}

	c.removeNode(n)
	c.addToFront(n)
}

func (c *lruCache) removeNode(n *node) {
	if n.prev != nil {
		n.prev.next = n.next
	} else {
		c.head = n.next
	}

	if n.next != nil {
		n.next.prev = n.prev
	} else {
		c.tail = n.prev
	}

	n.prev = nil
	n.next = nil
}
