package gouse

import (
	"errors"
	"sync"
	"time"

	"github.com/patrickmn/go-cache"
)

/* Local cache */

type ILocalCache struct {
	Set  map[string]string
	Lock sync.RWMutex
}

func NewLocalCache() *ILocalCache {
	return &ILocalCache{
		Set:  make(map[string]string),
		Lock: sync.RWMutex{},
	}
}

func (c *ILocalCache) GetLocalCache(key string) (string, error) {
	c.Lock.RLock()
	defer c.Lock.RUnlock()
	if c.Set == nil {
		return "", errors.New("set cache is not initialized")
	}
	return c.Set[key], nil
}

func (c *ILocalCache) SetLocalCache(key, value string) {
	c.Lock.Lock()
	defer c.Lock.Unlock()
	c.Set[key] = value
}

func (c *ILocalCache) DelLocalCache(key string) {
	c.Lock.Lock()
	defer c.Lock.Unlock()
	delete(c.Set, key)
}

func (c *ILocalCache) FlushLocalCache() {
	c.Lock.Lock()
	defer c.Lock.Unlock()
	c.Set = map[string]string{}
}

func (c *ILocalCache) AllLocalCache() map[string]string {
	c.Lock.RLock()
	defer c.Lock.RUnlock()
	return c.Set
}

/* temp cache */

type ITmpCache struct {
	Expires time.Duration
	Set     *cache.Cache
}

func NewTmpCache(expires ...time.Duration) *ITmpCache {
	var expire time.Duration
	if len(expires) > 0 {
		expire = expires[0]
	} else {
		expire = 24 * time.Hour
	}
	return &ITmpCache{
		Expires: expire,
		Set:     cache.New(expire, expire),
	}
}

func (c *ITmpCache) GetTmpCache(cacheKey string) interface{} {
	if val, found := c.Set.Get(cacheKey); found {
		return val
	}
	return nil
}

func (c *ITmpCache) SetTmpCache(cacheKey string, value interface{}, expireTime time.Duration) {
	if expireTime == 0 {
		expireTime = cache.DefaultExpiration
	}

	c.Set.Set(cacheKey, value, expireTime)
}

func (c *ITmpCache) DelTmpCache(cacheKey string) {
	c.Set.Delete(cacheKey)
}

func (c *ITmpCache) FlushTmpCache() {
	c.Set.Flush()
}

func (c *ITmpCache) AllTmpCache() map[string]string {
	result := make(map[string]string)
	for k := range c.Set.Items() {
		result[k] = ToString(c.GetTmpCache(k))
	}
	return result
}
