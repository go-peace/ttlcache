package ttlcache

import (
	"sync"
	"time"
)

type Cache struct {
	sync.RWMutex
	ttl   time.Duration
	items map[string]*Item
}

func NewCache(ttl time.Duration) *Cache {
	cache := &Cache{
		ttl:   ttl,
		items: map[string]*Item{},
	}
	cache.startCleanupTimer()
	return cache
}

func (cache *Cache) Set(key string, data string) {
	cache.Lock()
	defer cache.Unlock()
	expiration := time.Now().Add(cache.ttl)
	item := &Item{data: data, expires: &expiration}
	cache.items[key] = item
}

func (cache *Cache) Get(key string) (data string, found bool) {
	cache.Lock()
	defer cache.Unlock()
	item, exists := cache.items[key]
	if !exists || item.expired() {
		return
	}
	item.touch(cache.ttl)
	return item.data, true
}

func (cache *Cache) startCleanupTimer() {
	duration := cache.ttl
	if duration < time.Second {
		duration = time.Second
	}
	ticker := time.NewTicker(duration)
	go func() {
		for range ticker.C {
			cache.cleanup()
		}
	}()
}

func (cache *Cache) cleanup() {
	cache.Lock()
	defer cache.Unlock()
	for key, item := range cache.items {
		if item.expired() {
			delete(cache.items, key)
		}
	}
}
