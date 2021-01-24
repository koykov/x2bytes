# XToBytes

A collection of functions to convert to bytes. Allows converting anything to bytes.

## Usage

```go
var (
    buf []byte
    x float64 = 1234.56789
)
buf = x2bytes.ToBytes(buf, x) // []byte("1234.56789")
```

## Benchmark
```
BenchmarkToBytes-8   	  897764	      1296 ns/op	       0 B/op	       0 allocs/op
```
