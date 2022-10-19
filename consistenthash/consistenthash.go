package consistenthash

import (
	"hash/crc32"
	"sort"
)

// ConsistentHash 实际上，Key并不会加入到Map中，而只有Server会被加入。Key只用来查询最近的Server。Server可以多份加入。
type ConsistentHash struct {
	hashFunc func(data []byte) uint32
	ring     []int
}

func NewConsistentHash() *ConsistentHash {
	return &ConsistentHash{
		hashFunc: crc32.ChecksumIEEE,
		ring:     make([]int, 0),
	}
}

func (h *ConsistentHash) AddServer(server ...string) {
	for _, s := range server {
		h.ring = append(h.ring, int(h.hashFunc([]byte(s))))
	}
	sort.Ints(h.ring)
}

func (h *ConsistentHash) GetServer(key string) uint32 {
	if len(h.ring) == 0 {
		return 0
	}
	idx := sort.Search(len(h.ring), func(i int) bool {
		return h.ring[i] >= int(h.hashFunc([]byte(key)))
	})
	return uint32(idx % len(h.ring))
}
