package ttlcache

import (
	"sync"
	"time"
)

type item struct {
	sync.RWMutex
	data    string
	expires *time.Time
}

func (item *item) touch(duration time.Duration) {
	item.Lock()
	expiration := time.Now().Add(duration)
	item.expires = &expiration
	item.Unlock()
}

func (item *item) expired() bool {
	item.RLock()
	defer item.RUnlock()
	return item.expires == nil || item.expires.Before(time.Now())
}
