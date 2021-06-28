package ttlcache

import (
	"sync"
	"time"
)

// Cache is a in-memory thread safe map with ttl expiration.
type Cache interface {
	Set(string, string)
	Get(string) (string, bool)
	Cleanup()
	Len() int
}

// DefaultTTL will be applied if required ttl is less than 1 second when NewCache.
const DefaultTTL = time.Second

type mCache struct {
	sync.RWMutex
	ttl   time.Duration
	items map[string]*item
}

// NewCache returns a Cache with required ttl.
// If ttl is less than 1 second, a default ttl of 1 second will be applied.
func NewCache(ttl time.Duration) Cache {
	if ttl < DefaultTTL {
		ttl = DefaultTTL
	}
	cache := &mCache{
		ttl:   ttl,
		items: map[string]*item{},
	}
	cache.startCleanupTimer()
	return cache
}

// Set key/value to the cache
func (cache *mCache) Set(key string, value string) {
	cache.Lock()
	expiration := time.Now().Add(cache.ttl)
	item := &item{data: value, expires: &expiration}
	cache.items[key] = item
	cache.Unlock()
}

// Get from cache, returns value and a bool indicating whether key is found.
// Get will extend expire time if key exists
func (cache *mCache) Get(key string) (string, bool) {
	cache.RLock()
	item, exists := cache.items[key]
	cache.RUnlock()
	if !exists || item.expired() {
		return "", false
	}
	item.touch(cache.ttl)
	return item.data, true
}

// Len returns size of cached keys
func (cache *mCache) Len() int {
	cache.RLock()
	defer cache.RUnlock()
	return len(cache.items)
}

// Cleanup delete all expired keys from cache
func (cache *mCache) Cleanup() {
	for key, item := range cache.items {
		if item.expired() {
			cache.Lock()
			delete(cache.items, key)
			cache.Unlock()
		}
	}
}

func (cache *mCache) startCleanupTimer() {
	ticker := time.NewTicker(cache.ttl)
	go func() {
		for range ticker.C {
			cache.Cleanup()
		}
	}()
}
