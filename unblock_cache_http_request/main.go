package main

import (
	"fmt"
	"sync"
)

type Cache struct {
	cache   map[string]*Result
	cacheMu sync.Mutex
}

type Result struct {
	StatusCode int
	Status     string
	Body       []byte
	Err        error

	ready chan struct{} // is already request, but result not ready?
}

type RequestFunc func(url string) (*Result, error)

func NewCache() *Cache {
	return &Cache{
		cache: make(map[string]*Result, 8),
	}
}

// func (c *Cache) Get(key string, fn RequestFunc) (*Result, error) {
// 	c.cacheMu.Lock()
// 	defer c.cacheMu.Unlock()
// 	if res, ok := c.cache[key]; ok {
// 		return res, res.Err
// 	}

// 	fmt.Println("call request function: ", key)
// 	res, err := fn(key)
// 	c.cache[key] = res

// 	return res, err
// }

func (c *Cache) Get(key string, fn RequestFunc) (*Result, error) {
	c.cacheMu.Lock()
	res, ok := c.cache[key]
	if !ok {
		fmt.Println("call request function: ", key)

		newRes := &Result{ready: make(chan struct{})}
		c.cache[key] = newRes
		c.cacheMu.Unlock()

		res, err := fn(key)
		res.Err = err

		// *(newRes) = *res
		newRes.Status = res.Status
		newRes.StatusCode = res.StatusCode
		newRes.Body = res.Body
		newRes.Err = res.Err

		close(newRes.ready)

		return newRes, newRes.Err
	}

	c.cacheMu.Unlock()
	<-res.ready

	return res, res.Err
}

func main() {
}
