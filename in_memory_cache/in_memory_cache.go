package cache

import "fmt"

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
func (cache *Cache) Get(key string) interface{} {
	if _, ok := cache.keys[key]; ok {
		return cache.keys[key]
	}
	fmt.Printf("Key '%s' does not exist in memory cache", key)
	return nil
}

/*
Allows to delete key from memory cache
*/
func (cache *Cache) Delete(key string) {
	if _, ok := cache.keys[key]; ok {
		delete(cache.keys, key)
		fmt.Printf("Key '%s' successfully deleted from the memory cache", key)
	} else {
		fmt.Printf("Key '%s' does not exist in memory cache", key)
	}
}
