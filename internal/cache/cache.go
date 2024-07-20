package cache

import (
	"errors"
	"sync"
)

var (
	ErrNoRecord     = errors.New("models: no matching record found")
	ErrRecordExists = errors.New("models: such record already exists")
)

type Cache struct {
	storage *sync.Map
}

// Set a value in the cache
func (c *Cache) Insert(key string, value interface{}) error {
	if _, err := c.Get(key); err == nil {
		return ErrRecordExists
	}

	c.storage.Store(key, value)

	return nil
}

// Get a value by key from the cache
func (c *Cache) Get(key string) (interface{}, error) {
	val, ok := c.storage.Load(key)
	if !ok {
		return nil, ErrNoRecord
	}

	return val, nil
}

// Delete a value by key from the cache
func (c *Cache) Delete(key string) error {
	if _, err := c.Get(key); err != nil {
		return err
	}

	c.storage.Delete(key)

	return nil
}

func (c *Cache) Update(key string, value interface{}) error {
	if _, err := c.Get(key); err != nil {
		return err
	}

	c.storage.Swap(key, value)

	return nil
}

func (c *Cache) GetAll() []interface{} {
	var arr []interface{}

	f := func(key string, value interface{}) bool {
		arr = append(arr, value)

		return true
	}

	c.storage.Range(f)

	return arr
}
