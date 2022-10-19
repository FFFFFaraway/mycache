package server

import (
	"fmt"
	"log"
	"net/http"
)

type Handler struct {
	addr string
	s    *Server
}

func New(addr string, cacheSize int, callback func(string) any) *Handler {
	return &Handler{
		addr: addr,
		s:    NewServer(cacheSize, callback),
	}
}

// Log info with server name
func (p *Handler) Log(format string, v ...interface{}) {
	log.Printf("[Server %s] %s", p.addr, fmt.Sprintf(format, v...))
}

func (p *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p.Log("%s %s", r.Method, r.URL.Path)
	key := r.URL.Path[1:]
	res := p.s.Get(key)
	if res == nil {
		log.Printf("%s not exist\n", key)
		w.Write([]byte(key + " not exist"))
		return
	}
	w.Write([]byte(res.(string)))
}
