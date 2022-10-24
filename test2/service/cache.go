package service

import (
	"encoding/json"
	"fmt"
)

// NOTE: 추후 tenant까지 분리가 필요하면 그 때 map 작업
func NewCacheMap[T any](coll string) CacheMap[T] {
	return CacheMap[T]{
		collection: coll,
		data:       make(map[string]T),
	}
}

type CacheMap[T any] struct {
	collection string
	data       map[string]T
}

func (c CacheMap[T]) onInsert(id string, data interface{}) error {
	insertData, err := c.decode(data)
	if err != nil {
		return err
	}
	_, ok := c.Get(id)
	if !ok {
		c.data[id] = insertData
	}
	return nil
}

func (c CacheMap[T]) onUpdate(id string, data interface{}) error {
	updateData, err := c.decode(data)
	if err != nil {
		return err
	}
	_, ok := c.Get(id)
	if !ok {
		return fmt.Errorf("key dose not exist: %v", id)
	}
	c.data[id] = updateData
	return nil
}

func (c CacheMap[T]) onDelete(id string) {
	delete(c.data, id)
}

func (c CacheMap[T]) decode(data interface{}) (T, error) {
	var t T
	b, err := json.Marshal(data)
	if err != nil {
		return t, fmt.Errorf("fail to marshal: %v", data)
	}
	json.Unmarshal(b, &t)
	if err != nil {
		return t, fmt.Errorf("fail to unmarshal: %v", string(b))
	}
	return t, nil
}

func (c CacheMap[T]) Get(key string) (T, bool) {
	v, ok := c.data[key]
	return v, ok
}
