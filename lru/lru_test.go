package lru

import (
	"testing"
)

func TestLRU(t *testing.T) {
	c := MakeLRU(2)
	c.Add("1", 1) // 缓存是 {1=1}
	c.Add("2", 2) // 缓存是 {1=1, 2=2}
	if c.Get("1") != 1 {
		t.Fatal("Failed")
	}

	c.Add("3", 3) // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
	if c.Get("2") != nil {
		t.Fatal("Failed") // 返回 nil (未找到)
	}
	c.Add("4", 4) // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
	if c.Get("1") != nil {
		t.Fatal("Failed") // 返回 nil (未找到)
	}
	if c.Get("3") != 3 {
		t.Fatal("Failed") // 返回 3
	}
	if c.Get("4") != 4 {
		t.Fatal("Failed") // 返回 4
	}
}
