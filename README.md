# ttlcache
[![Go Report Card](https://goreportcard.com/badge/github.com/go-peace/ttlcache)](https://goreportcard.com/report/github.com/go-peace/ttlcache)
[![Go](https://github.com/go-peace/ttlcache/actions/workflows/go.yml/badge.svg)](https://github.com/go-peace/ttlcache/actions/workflows/go.yml)
[![golangci-lint](https://github.com/go-peace/ttlcache/actions/workflows/golangci-lint.yml/badge.svg)](https://github.com/go-peace/ttlcache/actions/workflows/golangci-lint.yml)

a thread-safe and fast cache to store string with expire ttl

## Usage



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
