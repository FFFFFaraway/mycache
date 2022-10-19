package consistenthash

import "hash/crc32"

type IHash interface {
	GetServer(key string) uint32
}

// Hash Mod N会导致机器数量变化时，很多Cache需要重新获取
type Hash struct {
	hashFunc func(data []byte) uint32
	n        uint32
}

func NewHash(n int) *Hash {
	return &Hash{
		hashFunc: crc32.ChecksumIEEE,
		n:        uint32(n),
	}
}

func (h *Hash) GetServer(key string) uint32 {
	return h.hashFunc([]byte(key)) % h.n
}
