package lru

import (
	"mycache/deque"
	"sync"
)

type Entry struct {
	k string
	v any
}

type LRU struct {
	mu  sync.Mutex
	cap int
	q   deque.Deque
	m   map[string]*deque.Node
}

func MakeLRU(cap int) *LRU {
	return &LRU{
		cap: cap,
		m:   make(map[string]*deque.Node),
	}
}

func (c *LRU) Add(k string, v any) {
	c.mu.Lock()
	defer c.mu.Unlock()
	n, exist := c.m[k]
	if exist {
		n.Val.(*Entry).v = v
		c.q.Remove(n)
		c.q.Append(n)
		return
	}
	if c.q.Len() == c.cap {
		n := c.q.Pop()
		delete(c.m, n.Val.(*Entry).k)
	}
	n = &deque.Node{Val: &Entry{k: k, v: v}}
	c.q.Append(n)
	c.m[k] = n
}

func (c *LRU) Get(k string) any {
	c.mu.Lock()
	defer c.mu.Unlock()
	n, exist := c.m[k]
	if !exist {
		return nil
	}
	c.q.Remove(n)
	c.q.Append(n)
	return n.Val.(*Entry).v
}
