package clerk

import (
	"fmt"
	"io/ioutil"
	"log"
	"mycache/consistenthash"
	"net/http"
	"net/url"
)

type Clerk struct {
	addr       string
	serverAddr []string
	hash       *consistenthash.ConsistentHash
}

func New(addr string, servers []string) *Clerk {
	hash := consistenthash.NewConsistentHash()
	hash.AddServer(servers...)
	return &Clerk{
		addr:       addr,
		serverAddr: servers,
		hash:       hash,
	}
}

// Log info with server name
func (c *Clerk) Log(format string, v ...interface{}) {
	log.Printf("[Server %s] %s", c.addr, fmt.Sprintf(format, v...))
}

func (c *Clerk) key2idx(key string) int {
	return int(c.hash.GetServer(key))
}

func (c *Clerk) Get(key string) ([]byte, error) {
	u := fmt.Sprintf(
		"http://%v/%v",
		c.serverAddr[c.key2idx(key)],
		url.QueryEscape(key),
	)
	res, err := http.Get(u)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("server returned: %v", res.Status)
	}

	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response body: %v", err)
	}

	return bytes, nil
}

func (c *Clerk) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c.Log("%s %s", r.Method, r.URL.Path)
	key := r.URL.Path[1:]

	res, err := c.Get(key)
	if err != nil {
		log.Printf("error: %s\n", err)
		w.Write([]byte("error: " + err.Error()))
		return
	}
	if res == nil {
		log.Printf("%s not exist\n", key)
		w.Write([]byte(key + " not exist"))
		return
	}
	w.Write(res)
}
