package deque

import "testing"

func TestDeque(t *testing.T) {
	q := Deque{}
	q.Append(&Node{Val: 1})
	q.Append(&Node{Val: 2})
	q.Append(&Node{Val: 3})
	q.AppendLeft(&Node{Val: 4})
	q.AppendLeft(&Node{Val: 5})
	q.AppendLeft(&Node{Val: 6})

	// 6 5 4 1 2 3
	if q.Pop().Val.(int) != 6 {
		t.Fatal("Failed")
	}
	if q.PopRight().Val.(int) != 3 {
		t.Fatal("Failed")
	}
	if q.Pop().Val.(int) != 5 {
		t.Fatal("Failed")
	}
	if q.PopRight().Val.(int) != 2 {
		t.Fatal("Failed")
	}
	if q.Pop().Val.(int) != 4 {
		t.Fatal("Failed")
	}
	if q.PopRight().Val.(int) != 1 {
		t.Fatal("Failed")
	}
	if q.PopRight() != nil {
		t.Fatal("Failed")
	}
}
