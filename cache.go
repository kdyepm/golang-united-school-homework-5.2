package cache

import (
	"time"
)

// func main() {
// 	cache := NewCache()
// 	fmt.Println(cache.Keys())
// 	cache.Put("somekey1", "someval1")
// 	cache.Put("somekey2", "someval2")
// 	fmt.Println(cache.Keys())
// 	fmt.Println(cache.Get("somekey2"))
// 	tthen := time.Now().Add(7 * time.Second)
// 	fmt.Println(cache.Keys())
// 	cache.PutTill("somekey3", "someval3", tthen)
// 	fmt.Println(cache.Keys())
// 	time.Sleep(2 * time.Second)
// 	fmt.Println(cache.Keys())
// 	fmt.Println(cache.Get("somekey2"))
// 	time.Sleep(7 * time.Second)
// 	fmt.Println(cache.Keys())
// }

type Cache struct {
	kvs       map[string]string
	deadlines map[string]time.Time
}

func (c *Cache) CleanUp(key string, deadline time.Time) {
	val, ok := c.deadlines[key]
	if !ok {
		return
	}
	if deadline.After(val) {
		delete(c.deadlines, key)
		delete(c.kvs, key)
	}
}

func NewCache() Cache {
	return Cache{make(map[string]string), make(map[string]time.Time)}
}

func (c *Cache) Get(key string) (string, bool) {
	tnow := time.Now()
	c.CleanUp(key, tnow)
	value, ok := c.kvs[key]
	return value, ok
}

func (c *Cache) Put(key, value string) {
	c.kvs[key] = value
}

func (c *Cache) Keys() []string {
	tnow := time.Now()
	keys := make([]string, 0, len(c.kvs))
	for k := range c.kvs {
		c.CleanUp(k, tnow)
	}
	for k := range c.kvs {
		keys = append(keys, k)

	}
	return keys
}

func (c *Cache) PutTill(key, value string, deadline time.Time) {
	c.kvs[key] = value
	c.deadlines[key] = deadline
}
