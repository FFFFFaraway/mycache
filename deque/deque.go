package deque

import "fmt"

type Node struct {
	Val   any
	left  *Node
	right *Node
}

func (n *Node) iso() {
	n.left = nil
	n.right = nil
}

type Deque struct {
	head *Node
	tail *Node
	len  int
}

func (q *Deque) Len() int {
	return q.len
}

func (q *Deque) out() {
	for i := q.head; i != nil; i = i.right {
		fmt.Printf("%v ", i.Val)
	}
	fmt.Println("Len: ", q.len)
}

func (q *Deque) Remove(n *Node) {
	if n == nil {
		return
	}
	defer func() { n.iso(); q.len -= 1 }()
	if n == q.head && n == q.tail {
		q.head = nil
		q.tail = nil
		return
	}
	if n == q.head {
		q.head = q.head.right
		return
	}
	if n == q.tail {
		q.tail = q.tail.left
		return
	}
	n.left.right, n.right.left = n.right, n.left
}

func (q *Deque) Append(n *Node) {
	q.len += 1
	if q.head == nil {
		q.head = n
		q.tail = n
		return
	}
	q.tail.right = n
	n.left = q.tail
	q.tail = n
}

func (q *Deque) AppendLeft(n *Node) {
	q.len += 1
	if q.head == nil {
		q.head = n
		q.tail = n
		return
	}
	q.head.left = n
	n.right = q.head
	q.head = n
}

func (q *Deque) Pop() *Node {
	res := q.head
	q.Remove(q.head)
	return res
}

func (q *Deque) PopRight() *Node {
	res := q.tail
	q.Remove(q.tail)
	return res
}
