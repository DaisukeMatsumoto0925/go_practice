package main

import (
	"fmt"
	"sync"
	"time"
)

// func main() {
// 	c := make(map[string]int)

// 	go func() {
// 		for i := 0; i < 1000; i++ {
// 			c["somekey"] += 1
// 		}
// 	}()

// 	go func() {
// 		for i := 0; i < 1000; i++ {
// 			c["somekey"] += 1
// 		}
// 	}()

// 	time.Sleep(time.Second)
// 	fmt.Println(c, c["somekey"])
// }

type safeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

func (c *safeCounter) inc(key string) {
	c.mux.Lock()
	c.v[key]++
	c.mux.Unlock()
}

func (c *safeCounter) value(key string) int {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.v[key]
}

func main() {
	c := safeCounter{v: make(map[string]int)}
	go func() {
		for i := 0; i < 1000; i++ {
			c.inc("somekey")
			c.inc("somekey2")
		}
	}()
	go func() {
		for i := 0; i < 1000; i++ {
			c.inc("somekey")
			c.inc("somekey3")
		}
	}()

	time.Sleep(time.Second)
	fmt.Println(c, c.value("somekey"))
}
