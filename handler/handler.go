package handler

import (
	"fmt"
	"log"
	"mycache/lru"
	"net/http"
)

type Handler struct {
	addr     string
	cache    *lru.LRU
	callback func(string) any
}

func New(addr string, cacheSize int, callback func(string) any) *Handler {
	return &Handler{
		addr:     addr,
		cache:    lru.MakeLRU(cacheSize),
		callback: callback,
	}
}

// Log info with server name
func (p *Handler) Log(format string, v ...interface{}) {
	log.Printf("[Server %s] %s", p.addr, fmt.Sprintf(format, v...))
}

func (p *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p.Log("%s %s", r.Method, r.URL.Path)
	key := r.URL.Path[1:]
	res := p.cache.Get(key)
	if res == nil {
		dbRes := p.callback(key)
		if dbRes == nil {
			log.Printf("%s not exist\n", key)
			return
		}
		w.Write([]byte(dbRes.(string)))
		p.cache.Add(key, dbRes)
		return
	}
	log.Println("[GeeCache] hit")
	w.Write([]byte(res.(string)))
}
