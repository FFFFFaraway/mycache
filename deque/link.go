package deque

type OneNode struct {
	val  int
	next *OneNode
}

type Link struct {
	head *OneNode
	tail *OneNode
}

func MakeLink() *Link {
	return &Link{}
}

func (l *Link) Append(x int) {
	n := &OneNode{val: x}
	if l.head == nil {
		l.head = n
		l.tail = n
		return
	}
	l.tail.next = n
	l.tail = n
}
