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

func (c *Cache) Get(key string) (string, bool) {
	data, exist := c.store[key]
	if !exist {
		return "", false
	} else if !time.Now().Before(data.expiredAt) {
		delete(c.store, key)
		return "", false
	}
	return data.value, exist
}

func (c *Cache) Put(key, value string) {
	c.store[key] = data{
		value:     value,
		expiredAt: time.Now().AddDate(100, 0, 0),
	}
}

func (c *Cache) Keys() []string {
	keys := make([]string, 0)
	for key, data := range c.store {
		if time.Now().Before(data.expiredAt) {
			keys = append(keys, key)
		}
	}
	return keys
}

func (c *Cache) PutTill(key, value string, deadline time.Time) {
	c.store[key] = data{
		value:     value,
		expiredAt: deadline,
	}
}
