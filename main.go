package main

/*
   $ curl http://localhost:9999/Tom
   630

   $ curl http://localhost:9999/kkk
   kkk not exist
*/

import (
	"log"
	"mycache/handler"
	"net/http"
)

var db = map[string]string{
	"Tom":  "630",
	"Jack": "589",
	"Sam":  "567",
}

func main() {
	addr := "localhost:9999"
	h := handler.New(addr, 2, func(key string) any {
		log.Println("[SlowDB] search key", key)
		v, exist := db[key]
		if !exist {
			return nil
		}
		return v
	})
	log.Println("cache is running at", addr)
	log.Fatal(http.ListenAndServe(addr, h))
}
