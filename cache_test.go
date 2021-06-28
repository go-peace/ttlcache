package ttlcache

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestCleanupLatency(t *testing.T) {
	t.Run("100kv", func(t *testing.T) {
		cache := NewCache(0)
		for _, kv := range genKV(100) {
			cache.Set(kv[0], kv[1])
		}
		time.Sleep(1100 * time.Millisecond)
		now := time.Now()
		cache.Cleanup()
		fmt.Printf("delete 100kv\tlatency: %v\n", time.Since(now))
		if cache.Len() != 0 {
			t.Error("cache should be cleanup")
		}
	})
	t.Run("1000kv", func(t *testing.T) {
		cache := NewCache(0)
		for _, kv := range genKV(1000) {
			cache.Set(kv[0], kv[1])
		}
		time.Sleep(1100 * time.Millisecond)
		now := time.Now()
		cache.Cleanup()
		fmt.Printf("delete 1000kv\tlatency: %v\n", time.Since(now))
		if cache.Len() != 0 {
			t.Error("cache should be cleanup")
		}
	})
	t.Run("10000kv", func(t *testing.T) {
		cache := NewCache(0)
		for _, kv := range genKV(10000) {
			cache.Set(kv[0], kv[1])
		}
		time.Sleep(1100 * time.Millisecond)
		now := time.Now()
		cache.Cleanup()
		fmt.Printf("delete 10000kv\tlatency: %v\n", time.Since(now))
		if cache.Len() != 0 {
			t.Error("cache should be cleanup")
		}
	})
	t.Run("100000kv", func(t *testing.T) {
		cache := NewCache(0)
		for _, kv := range genKV(100000) {
			cache.Set(kv[0], kv[1])
		}
		time.Sleep(1100 * time.Millisecond)
		now := time.Now()
		cache.Cleanup()
		fmt.Printf("delete 100000kv\tlatency: %v\n", time.Since(now))
		if cache.Len() != 0 {
			t.Error("cache should be cleanup")
		}
	})
}

func BenchmarkCache(b *testing.B) {
	kvs := genKV(100000)
	cache := NewCache(time.Hour)

	// -------- 1000kv ----------
	n := 1000
	for i := 0; i < n; i++ {
		cache.Set(kvs[i][0], kvs[i][1])
	}
	if cache.Len() != n {
		b.FailNow()
	}
	b.Run("1000kv Set", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			cache.Set(kvs[0][0], kvs[0][1])
		}
	})
	b.Run("1000kv Get", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			cache.Get(kvs[0][0])
		}
	})

	fmt.Println()
	// ---------- 10000 kv --------------
	n = 10000
	for i := 0; i < n; i++ {
		cache.Set(kvs[i][0], kvs[i][1])
	}
	if cache.Len() != n {
		b.FailNow()
	}
	b.Run("10000kv Set", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			cache.Set(kvs[0][0], kvs[0][1])
		}
	})
	b.Run("10000kv Get", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			cache.Get(kvs[0][0])
		}
	})

	fmt.Println()
	// ---------- 100000 kv ------------
	n = 100000
	for i := 0; i < n; i++ {
		cache.Set(kvs[i][0], kvs[i][1])
	}
	if cache.Len() != n {
		b.FailNow()
	}
	b.Run("100000kv Set", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			cache.Set(kvs[0][0], kvs[0][1])
		}
	})
	b.Run("100000kv Get", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			cache.Get(kvs[0][0])
		}
	})

}

func genKV(n int) [][2]string {
	kvs := make([][2]string, n)
	for i := 0; i < n; i++ {
		kvs[i][0], kvs[i][1] = randStr(20), randStr(15)
	}
	return kvs
}

func randStr(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
