package cache

import "time"

type data struct {
	value     string
	expiredAt time.Time
}

type Cache struct {
	store map[string]data
}

func NewCache() Cache {
	return Cache{
		store: make(map[string]data),
	}
}

func (c Cache) Get(key string) (string, bool) {
	data, ok := c.store[key]
	if !ok || time.Now().After(data.expiredAt) {
		return "", false
	}
	return data.value, ok
}

func (c *Cache) Put(key, value string) {
	c.store[key] = data{value: value, expiredAt: time.Unix(1<<63-1, 0)}
}

func (c Cache) Keys() []string {
	keys := make([]string, 0)
	for key, data := range c.store {
		if !time.Now().After(data.expiredAt) {
			keys = append(keys, key)
		}
	}
	return keys
}

func (c Cache) PutTill(key, value string, deadline time.Time) {
	c.store[key] = data{
		value:     value,
		expiredAt: deadline,
	}
	return
}
