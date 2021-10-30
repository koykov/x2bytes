# X2Bytes

A collection of functions to convert anything to bytes.

## Usage

```go
var (
    buf []byte
    x float64 = 1234.56789
)
buf = x2bytes.ToBytes(buf, x) // []byte("1234.56789")
```

Also file [x2bytes_builtin.go](x2bytes_builtin.go) contains functions to convert builtin types to bytes. You may use
them separately.

## Benchmarks
```
BenchmarkToBytes/bytes-8         	59374212	       20.14 ns/op	       0 B/op	       0 allocs/op
BenchmarkToBytes/*bytes-8        	58059764	       20.05 ns/op	       0 B/op	       0 allocs/op
BenchmarkToBytes/string-8        	34148546	       36.04 ns/op	       0 B/op	       0 allocs/op
BenchmarkToBytes/*string-8       	36105606	       33.43 ns/op	       0 B/op	       0 allocs/op
BenchmarkToBytes/bool-8          	32572970	       36.75 ns/op	       0 B/op	       0 allocs/op
BenchmarkToBytes/*bool-8         	32717931	       37.17 ns/op	       0 B/op	       0 allocs/op
BenchmarkToBytes/int-8           	18412017	       65.02 ns/op	       0 B/op	       0 allocs/op
BenchmarkToBytes/*int-8          	18307768	       66.23 ns/op	       0 B/op	       0 allocs/op
BenchmarkToBytes/int8-8          	18608317	       64.10 ns/op	       0 B/op	       0 allocs/op
BenchmarkToBytes/*int8-8         	18653878	       67.45 ns/op	       0 B/op	       0 allocs/op
BenchmarkToBytes/int16-8         	18617209	       64.26 ns/op	       0 B/op	       0 allocs/op
BenchmarkToBytes/*int16-8        	18493111	       64.27 ns/op	       0 B/op	       0 allocs/op
BenchmarkToBytes/int32-8         	18633570	       64.49 ns/op	       0 B/op	       0 allocs/op
BenchmarkToBytes/*int32-8        	18608048	       69.45 ns/op	       0 B/op	       0 allocs/op
BenchmarkToBytes/int64-8         	18161239	       64.42 ns/op	       0 B/op	       0 allocs/op
BenchmarkToBytes/*int64-8        	18435315	       67.21 ns/op	       0 B/op	       0 allocs/op
BenchmarkToBytes/uint-8          	16034138	       74.53 ns/op	       0 B/op	       0 allocs/op
BenchmarkToBytes/*uint-8         	16225942	       77.62 ns/op	       0 B/op	       0 allocs/op
BenchmarkToBytes/uint8-8         	15852429	       75.67 ns/op	       0 B/op	       0 allocs/op
BenchmarkToBytes/*uint8-8        	16295343	       78.85 ns/op	       0 B/op	       0 allocs/op
BenchmarkToBytes/uint16-8        	15533521	       75.82 ns/op	       0 B/op	       0 allocs/op
BenchmarkToBytes/*uint16-8       	15797869	       76.09 ns/op	       0 B/op	       0 allocs/op
BenchmarkToBytes/uint32-8        	16202493	       73.83 ns/op	       0 B/op	       0 allocs/op
BenchmarkToBytes/*uint32-8       	15437512	       76.56 ns/op	       0 B/op	       0 allocs/op
BenchmarkToBytes/uint64-8        	16051436	       75.10 ns/op	       0 B/op	       0 allocs/op
BenchmarkToBytes/*uint64-8       	16192809	       75.66 ns/op	       0 B/op	       0 allocs/op
BenchmarkToBytes/float32-8       	 4876089	       248.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkToBytes/*float32-8      	 4848788	       241.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkToBytes/float64-8       	 5956798	       206.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkToBytes/*float64-8      	 5938024	       199.3 ns/op	       0 B/op	       0 allocs/op
```
