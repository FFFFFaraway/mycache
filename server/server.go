package server

import (
	"log"
	"mycache/lru"
)

type Server struct {
	cache    *lru.LRU
	callback func(string) any
}

func NewServer(cacheSize int, callback func(string) any) *Server {
	return &Server{
		cache:    lru.MakeLRU(cacheSize),
		callback: callback,
	}
}

func (s *Server) Get(key string) any {
	res := s.cache.Get(key)
	if res == nil {
		dbRes := s.callback(key)
		if dbRes == nil {
			return nil
		}
		s.cache.Add(key, dbRes)
		return dbRes
	}
	log.Println("[GeeCache] hit")
	return res
}
