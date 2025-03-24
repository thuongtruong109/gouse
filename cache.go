package gouse

import (
	"errors"
	"sync"
	"time"

	"github.com/patrickmn/go-cache"
)

type ILCache struct {
	Set  map[string]string
	Lock sync.RWMutex
}

func NewLCache() *ILCache {
	return &ILCache{
		Set:  make(map[string]string),
		Lock: sync.RWMutex{},
	}
}

func (c *ILCache) GetLCache(key string) (string, error) {
	c.Lock.RLock()
	defer c.Lock.RUnlock()
	if c.Set == nil {
		return "", errors.New("set cache is not initialized")
	}
	return c.Set[key], nil
}

func (c *ILCache) SetLCache(key, value string) {
	c.Lock.Lock()
	defer c.Lock.Unlock()
	c.Set[key] = value
}

func (c *ILCache) DelLCache(key string) {
	c.Lock.Lock()
	defer c.Lock.Unlock()
	delete(c.Set, key)
}

func (c *ILCache) FlushLCache() {
	c.Lock.Lock()
	defer c.Lock.Unlock()
	c.Set = map[string]string{}
}

func (c *ILCache) AllLCache() map[string]string {
	c.Lock.RLock()
	defer c.Lock.RUnlock()
	return c.Set
}

type ITCache struct {
	Expires time.Duration
	Set     *cache.Cache
}

func NewTCache(expires ...time.Duration) *ITCache {
	var expire time.Duration
	if len(expires) > 0 {
		expire = expires[0]
	} else {
		expire = 24 * time.Hour
	}
	return &ITCache{
		Expires: expire,
		Set:     cache.New(expire, expire),
	}
}

func (c *ITCache) GetTCache(cacheKey string) interface{} {
	if val, found := c.Set.Get(cacheKey); found {
		return val
	}
	return nil
}

func (c *ITCache) SetTCache(cacheKey string, value interface{}, expireTime time.Duration) {
	if expireTime == 0 {
		expireTime = cache.DefaultExpiration
	}

	c.Set.Set(cacheKey, value, expireTime)
}

func (c *ITCache) DelTCache(cacheKey string) {
	c.Set.Delete(cacheKey)
}

func (c *ITCache) FlushTCache() {
	c.Set.Flush()
}

func (c *ITCache) AllTCache() map[string]string {
	result := make(map[string]string)
	for k := range c.Set.Items() {
		result[k] = ToStr(c.GetTCache(k))
	}
	return result
}
