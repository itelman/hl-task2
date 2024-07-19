package cache

import (
	"sync"
)

type Cache struct {
	store sync.Map
}

// Set a value in the cache
func (c *Cache) Insert(key string, value interface{}) {
	c.store.Store(key, value)
}

// Get a value by key from the cache
func (c *Cache) Get(key string) (interface{}, bool) {
	val, ok := c.store.Load(key)
	return val, ok
}

// Delete a value by key from the cache
func (c *Cache) Delete(key string) {
	c.store.Delete(key)
}

/*func (c *Cache) GetAll(key string) {
	c.store.Range(c.store.Load())
}*/
