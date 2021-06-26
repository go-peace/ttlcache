package ttlcache

import (
	"sync"
	"time"
)

type Item struct {
	sync.RWMutex
	data    string
	expires *time.Time
}

func (item *Item) touch(duration time.Duration) {
	item.Lock()
	defer item.Unlock()
	expiration := time.Now().Add(duration)
	item.expires = &expiration
}

func (item *Item) expired() bool {
	item.RLock()
	defer item.RUnlock()
	return item.expires == nil || item.expires.Before(time.Now())
}
