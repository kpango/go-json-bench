# This repository is Archived, no longer maintained.


# Golang JSON Parse Benchmark
golang json Unmarshal VS Decode Benchmark

## Description
There is Unmarshal and Decoder in golang's encoding/json package, and both can be used when parsing JSON data.
In this repository, we compared the io.Reader parsing performance of both by performing a simple benchmark.

## Code Example
### json.Unmarshal
```go
	var target JSON
	body, err := ioutil.ReadAll(reader)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, &target)
```

### json.Decoder
```go
	var target JSON
	return json.NewDecoder(reader).Decode(&target)
```

## Motivation
In golang, compare them to make a performance-oriented choice when parsing json

## Installation
```
go get github.com/kpango/go-json-bench
```

## Usage
### Benchmark
```
make bench
```
### Profiler
```
make profile
```
### Cleaning
```
make clean
```


## Benchmark Result
```
go test -count=10 -run=NONE -bench . -benchmem -o pprof/test.bin -cpuprofile pprof/cpu.out -memprofile pprof/mem.out
goos: darwin
goarch: amd64
pkg: github.com/kpango/json-bench
BenchmarkJSONUnmarshal-8           	  500000	      2925 ns/op	    2640 B/op	      10 allocs/op
BenchmarkJSONUnmarshal-8           	  500000	      2958 ns/op	    2640 B/op	      10 allocs/op
BenchmarkJSONUnmarshal-8           	  500000	      3052 ns/op	    2640 B/op	      10 allocs/op
BenchmarkJSONUnmarshal-8           	  500000	      2959 ns/op	    2640 B/op	      10 allocs/op
BenchmarkJSONUnmarshal-8           	  500000	      2976 ns/op	    2640 B/op	      10 allocs/op
BenchmarkJSONUnmarshal-8           	  500000	      2942 ns/op	    2640 B/op	      10 allocs/op
BenchmarkJSONUnmarshal-8           	  500000	      2977 ns/op	    2640 B/op	      10 allocs/op
BenchmarkJSONUnmarshal-8           	  500000	      2982 ns/op	    2640 B/op	      10 allocs/op
BenchmarkJSONUnmarshal-8           	  500000	      2936 ns/op	    2640 B/op	      10 allocs/op
BenchmarkJSONUnmarshal-8           	  500000	      2933 ns/op	    2640 B/op	      10 allocs/op
BenchmarkJSONDecode-8              	  500000	      2550 ns/op	    1192 B/op	       9 allocs/op
BenchmarkJSONDecode-8              	  500000	      2514 ns/op	    1192 B/op	       9 allocs/op
BenchmarkJSONDecode-8              	  500000	      2521 ns/op	    1192 B/op	       9 allocs/op
BenchmarkJSONDecode-8              	  500000	      2572 ns/op	    1192 B/op	       9 allocs/op
BenchmarkJSONDecode-8              	  500000	      2519 ns/op	    1192 B/op	       9 allocs/op
BenchmarkJSONDecode-8              	  500000	      2511 ns/op	    1192 B/op	       9 allocs/op
BenchmarkJSONDecode-8              	  500000	      2492 ns/op	    1192 B/op	       9 allocs/op
BenchmarkJSONDecode-8              	  500000	      2485 ns/op	    1192 B/op	       9 allocs/op
BenchmarkJSONDecode-8              	  500000	      2488 ns/op	    1192 B/op	       9 allocs/op
BenchmarkJSONDecode-8              	  500000	      2476 ns/op	    1192 B/op	       9 allocs/op
BenchmarkJSONUnmarshalParallel-8   	 1000000	      1194 ns/op	    2640 B/op	      10 allocs/op
BenchmarkJSONUnmarshalParallel-8   	 1000000	      1236 ns/op	    2640 B/op	      10 allocs/op
BenchmarkJSONUnmarshalParallel-8   	 1000000	      1243 ns/op	    2640 B/op	      10 allocs/op
BenchmarkJSONUnmarshalParallel-8   	 1000000	      1268 ns/op	    2640 B/op	      10 allocs/op
BenchmarkJSONUnmarshalParallel-8   	 1000000	      1251 ns/op	    2640 B/op	      10 allocs/op
BenchmarkJSONUnmarshalParallel-8   	 1000000	      1344 ns/op	    2640 B/op	      10 allocs/op
BenchmarkJSONUnmarshalParallel-8   	 1000000	      1360 ns/op	    2640 B/op	      10 allocs/op
BenchmarkJSONUnmarshalParallel-8   	 1000000	      1277 ns/op	    2640 B/op	      10 allocs/op
BenchmarkJSONUnmarshalParallel-8   	 1000000	      1264 ns/op	    2640 B/op	      10 allocs/op
BenchmarkJSONUnmarshalParallel-8   	 1000000	      1312 ns/op	    2640 B/op	      10 allocs/op
BenchmarkJSONDecodeParallel-8      	 2000000	       914 ns/op	    1192 B/op	       9 allocs/op
BenchmarkJSONDecodeParallel-8      	 2000000	       966 ns/op	    1192 B/op	       9 allocs/op
BenchmarkJSONDecodeParallel-8      	 2000000	       984 ns/op	    1192 B/op	       9 allocs/op
BenchmarkJSONDecodeParallel-8      	 2000000	       992 ns/op	    1192 B/op	       9 allocs/op
BenchmarkJSONDecodeParallel-8      	 2000000	      1062 ns/op	    1192 B/op	       9 allocs/op
BenchmarkJSONDecodeParallel-8      	 1000000	      1039 ns/op	    1192 B/op	       9 allocs/op
BenchmarkJSONDecodeParallel-8      	 1000000	      1021 ns/op	    1192 B/op	       9 allocs/op
BenchmarkJSONDecodeParallel-8      	 2000000	      1026 ns/op	    1192 B/op	       9 allocs/op
BenchmarkJSONDecodeParallel-8      	 1000000	      1016 ns/op	    1192 B/op	       9 allocs/op
BenchmarkJSONDecodeParallel-8      	 2000000	      1117 ns/op	    1192 B/op	       9 allocs/op
PASS
ok  	github.com/kpango/json-bench	65.207s
```
[full result](https://github.com/kpango/go-json-bench/blob/master/result.tsv)

## Conclusion
In the case of io.Reader's input data, the execution speed was slightly faster using json.Decoder, and the byte transfer amount was about half.
