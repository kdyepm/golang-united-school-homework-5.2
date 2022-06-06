package cache

import (
	"time"
)

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
