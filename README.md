# ttlcache
[![Go Report Card](https://goreportcard.com/badge/github.com/go-peace/ttlcache)](https://goreportcard.com/report/github.com/go-peace/ttlcache)
[![Go](https://github.com/go-peace/ttlcache/actions/workflows/go.yml/badge.svg)](https://github.com/go-peace/ttlcache/actions/workflows/go.yml)
[![golangci-lint](https://github.com/go-peace/ttlcache/actions/workflows/golangci-lint.yml/badge.svg)](https://github.com/go-peace/ttlcache/actions/workflows/golangci-lint.yml)

a thread-safe and fast cache to store string with expire ttl

---

## Usage
```go
package main

import (
	"fmt"
	"time"

	"github.com/go-peace/ttlcache"
)

func main() {
	// new a cache with ttl 1 second
	cache := ttlcache.NewCache(time.Second)

	// set value
	cache.Set("foo", "bar")

	// get value
	v, exist := cache.Get("foo")
	fmt.Printf("exist:%v\tvalue:%v\tcache size:%d\n", exist, v, cache.Len())

	// get value after expiration
	time.Sleep(2 * time.Second)
	v, exist = cache.Get("foo")
	fmt.Printf("exist:%v\tvalue:   %v\tcache size:%d\n", exist, v, cache.Len())
}
```
stdout:
```bash
exist:true	value:bar	cache size:1
exist:false	value:   	cache size:0
```



## Benchmark
```bash
➜  ttlcache git:(main) ✗ go test -bench=.
delete 100kv	latency: 8.209µs
delete 1000kv	latency: 91.125µs
delete 10000kv	latency: 947.833µs
delete 100000kv	latency: 12.819667ms

goos: darwin
goarch: arm64
pkg: github.com/go-peace/ttlcache
BenchmarkCache/1000kv_Set-8         	10195526	       114.3 ns/op
BenchmarkCache/1000kv_Get-8         	 7028764	       164.2 ns/op

BenchmarkCache/10000kv_Set-8        	10975851	       114.2 ns/op
BenchmarkCache/10000kv_Get-8        	 6717267	       168.5 ns/op

BenchmarkCache/100000kv_Set-8       	10080742	       120.0 ns/op
BenchmarkCache/100000kv_Get-8       	 6850396	       169.7 ns/op

PASS
ok  	github.com/go-peace/ttlcache	13.010s
```
