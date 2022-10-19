package consistenthash

import (
	"hash/crc32"
	"log"
	"sort"
	"strconv"
)

// MultiHash 实际上，Key并不会加入到Map中，而只有Server会被加入。Key只用来查询最近的Server。Server可以多份加入。
type MultiHash struct {
	hashFunc    func(data []byte) uint32
	ring        []int
	replicas    int
	hash2server map[int]string
}

func NewMultiHash(replicas int) *MultiHash {
	return &MultiHash{
		hashFunc:    crc32.ChecksumIEEE,
		ring:        make([]int, 0),
		hash2server: make(map[int]string),
		replicas:    replicas,
	}
}

func (h *MultiHash) AddServer(server ...string) {
	for _, s := range server {
		for i := 0; i < h.replicas; i++ {
			hash := int(h.hashFunc([]byte(strconv.Itoa(i) + s)))
			h.ring = append(h.ring, hash)
			h.hash2server[hash] = s
		}
	}
	sort.Ints(h.ring)
}

func (h *MultiHash) GetServer(key string) string {
	if len(h.ring) == 0 {
		log.Fatal("h.ring == 0!")
		return ""
	}
	idx := sort.Search(len(h.ring), func(i int) bool {
		return h.ring[i] >= int(h.hashFunc([]byte(key)))
	})
	serverHash := h.ring[idx%len(h.ring)]
	return h.hash2server[serverHash]
}
