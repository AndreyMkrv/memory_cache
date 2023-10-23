package cache

import (
	"errors"
)

type Cache struct {
	keys map[string]interface{}
}

func New() Cache {
	return Cache{
		keys: make(map[string]interface{}),
	}
}

/*
Allows to create new value in memory cache, where key is uniq string and value any interface
*/
func (cache *Cache) Set(key string, value interface{}) {
	cache.keys[key] = value
}

/*
Allows to get value from memory cache based on key
*/
func (cache *Cache) Get(key string) (interface{}, error) {
	if _, ok := cache.keys[key]; ok {
		return cache.keys[key], nil
	}
	return nil, errors.New(KEY_DOES_NOT_EXIST)
}

/*
Allows to delete key from memory cache
*/
func (cache *Cache) Delete(key string) error {
	if _, ok := cache.keys[key]; ok {
		delete(cache.keys, key)
		return nil
	} else {
		return errors.New(KEY_DOES_NOT_EXIST)
	}
}
