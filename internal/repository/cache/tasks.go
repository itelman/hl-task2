package cache

import (
	"errors"
	"sync"
	"todo-list/internal/service/timego"
	"todo-list/pkg/models"
)

var (
	ErrNoRecord     = errors.New("models: no matching record found")
	ErrRecordExists = errors.New("models: such record already exists")
)

type Cache struct {
	store sync.Map
}

// Set a value in the cache
func (c *Cache) Insert(key string, val models.Task) error {
	if _, err := c.Get(val.ID); err == nil {
		return ErrRecordExists
	}

	c.store.Store(val.ID, val)
	return nil
}

// Get a value by key from the cache
func (c *Cache) Get(key string) (models.Task, error) {
	val, ok := c.store.Load(key)
	if !ok {
		return models.Task{}, ErrNoRecord
	}

	return val.(models.Task), nil
}

func (c *Cache) Update(key string, val models.Task) error {
	if _, err := c.Get(key); err != nil {
		return err
	}

	c.store.Swap(key, val)

	return nil
}

// Delete a value by key from the cache
func (c *Cache) Delete(key string) error {
	if _, err := c.Get(key); err != nil {
		return err
	}

	c.store.Delete(key)

	return nil
}

func (c *Cache) GetAll() ([]*models.Task, error) {
	var arr []interface{}
	var tasks []*models.Task

	f := func(key, value any) bool {
		arr = append(arr, value)

		return true
	}

	c.store.Range(f)

	for _, val := range arr {
		task := val.(models.Task)
		tasks = append(tasks, &task)
	}

	for _, task := range tasks {
		isWeekend, err := timego.IsWeekend(task.ActiveAt)
		if err != nil {
			return nil, err
		}

		if isWeekend {
			task.MarkAsWeekend()
		}
	}

	return tasks, nil
}
