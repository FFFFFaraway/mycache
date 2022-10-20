package clerk

import (
	"fmt"
	"io/ioutil"
	"log"
	"mycache/consistenthash"
	"net/http"
	"net/url"
	"sync"
)

type CallResult struct {
	v   []byte
	err error
}

type Clerk struct {
	addr       string
	serverAddr []string
	hash       *consistenthash.MultiHash
	doingMu    sync.RWMutex
	doing      map[string]chan CallResult
}

func New(addr string, servers []string) *Clerk {
	hash := consistenthash.NewMultiHash(50)
	hash.AddServer(servers...)
	return &Clerk{
		addr:       addr,
		serverAddr: servers,
		hash:       hash,
		doing:      make(map[string]chan CallResult),
	}
}

// Log info with server name
func (c *Clerk) Log(format string, v ...interface{}) {
	log.Printf("[Server %s] %s", c.addr, fmt.Sprintf(format, v...))
}

func (c *Clerk) Get(key string) ([]byte, error) {
	u := fmt.Sprintf(
		"http://%v/%v",
		c.hash.GetServer(key),
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
	var res CallResult
	c.doingMu.RLock()
	ch, exist := c.doing[key]
	c.doingMu.RUnlock()
	if exist {
		res = <-ch
		c.Log("avoid call")
	} else {
		c.doingMu.Lock()
		c.doing[key] = make(chan CallResult)
		c.doingMu.Unlock()
		res.v, res.err = c.Get(key)

		c.doingMu.Lock()
	loop:
		for {
			select {
			case c.doing[key] <- res:
			default:
				break loop
			}
		}
		delete(c.doing, key)
		c.doingMu.Unlock()
	}
	if res.err != nil {
		log.Printf("error: %s\n", res.err)
		w.Write([]byte("error: " + res.err.Error()))
		return
	}
	if res.v == nil {
		log.Printf("%s not exist\n", key)
		w.Write([]byte(key + " not exist"))
		return
	}
	w.Write(res.v)
}
