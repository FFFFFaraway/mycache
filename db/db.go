package db

import (
	"log"
)

var db = map[string]string{
	"Tom":  "630",
	"Jack": "589",
	"Sam":  "567",
}

// Load 假设这里有一个数据库
func Load(key string) any {
	log.Println("[SlowDB] search key", key)
	if res, exist := db[key]; exist {
		return res
	} else {
		return nil
	}
}
