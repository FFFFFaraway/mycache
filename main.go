package main

import (
	"log"
	"mycache/clerk"
	"mycache/db"
	"mycache/server"
	"net/http"
)

func main() {
	addr := "localhost:9999"
	servers := []string{
		"localhost:8001",
		"localhost:8002",
		"localhost:8003",
	}
	for _, s := range servers {
		go func(addr string) {
			s := server.New(addr, 2, db.Load)
			log.Println("cache server is running at", addr)
			log.Fatal(http.ListenAndServe(addr, s))
		}(s)
	}

	c := clerk.New(addr, servers)
	log.Println("cache clerk is running at", addr)
	log.Fatal(http.ListenAndServe(addr, c))
}
