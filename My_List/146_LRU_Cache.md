# 146. LRU Cache

Design and implement a data structure for Least Recently Used (LRU) cache. It should support the following operations: get and put.

get(key) - Get the value (will always be positive) of the key if the key exists in the cache, otherwise return -1.
put(key, value) - Set or insert the value if the key is not already present. When the cache reached its capacity, it should invalidate the least recently used item before inserting a new item.

The cache is initialized with a positive capacity.

Follow up:
Could you do both operations in O(1) time complexity?

Example:

LRUCache cache = new LRUCache( 2 \/\* capacity \*\/ );

```golang
cache.put(1, 1);
cache.put(2, 2);
cache.get(1);       // returns 1
cache.put(3, 3);    // evicts key 2
cache.get(2);       // returns -1 (not found)
cache.put(4, 4);    // evicts key 1
cache.get(1);       // returns -1 (not found)
cache.get(3);       // returns 3
cache.get(4);       // returns 4
```

[146. LRU Cache](https://leetcode.com/problems/lru-cache/)

```golang
/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
 
type DLinkedList struct {
	val  int
	key  int
	prev *DLinkedList
	next *DLinkedList
}

type LRUCache struct {
	capacity int
	head     *DLinkedList
	tail     *DLinkedList
	cache    map[int]*DLinkedList
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*DLinkedList),
	}
}

func (c *LRUCache) Get(key int) int {
	l, ok := c.cache[key]
	if !ok {
		return -1
	}
	c.removeFromList(l)
	c.addToHead(l)
	return l.val
}

func (c *LRUCache) Put(key int, value int) {
	l, ok := c.cache[key]
	if !ok {
        // new key
		l = &DLinkedList{val: value, key: key}
		c.cache[key] = l
	} else {
        // update key
		l.val = value
		c.removeFromList(l)
	}

	c.addToHead(l)
	if len(c.cache) > c.capacity {
		l = c.tail
		if l != nil {
			c.removeFromList(l)
			delete(c.cache, l.key)
		}
	}
}

func (c *LRUCache) addToHead(l *DLinkedList) {
	l.prev = nil
	l.next = c.head
	if l.next != nil {
		l.next.prev = l
	}
	c.head = l
	if c.tail == nil {
		c.tail = l
	}
}

func (c *LRUCache) removeFromList(l *DLinkedList) {
	if l == c.head {
		c.head = l.next
	}

	if l == c.tail {
		c.tail = l.prev
	}

	if l.next != nil {
		l.next.prev = l.prev
	}

	if l.prev != nil {
		l.prev.next = l.next
	}
}
```
