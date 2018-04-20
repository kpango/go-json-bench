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
go tool pprof --svg pprof/test.bin pprof/mem.out > mem.svg
go tool pprof --svg pprof/test.bin pprof/cpu.out > cpu.svg
go test -count=10 -run=NONE -bench=BenchmarkJSONUnmarshal -benchmem -o pprof/test.bin -cpuprofile pprof/cpu-unmarshal.out -memprofile pprof/mem-unmarshal.out
goos: darwin
goarch: amd64
pkg: github.com/kpango/json-bench
BenchmarkJSONUnmarshal-8           	  500000	      3116 ns/op	    2640 B/op	      10 allocs/op
BenchmarkJSONUnmarshal-8           	  500000	      3022 ns/op	    2640 B/op	      10 allocs/op
BenchmarkJSONUnmarshal-8           	  500000	      3041 ns/op	    2640 B/op	      10 allocs/op
BenchmarkJSONUnmarshal-8           	  500000	      3102 ns/op	    2640 B/op	      10 allocs/op
BenchmarkJSONUnmarshal-8           	  500000	      3002 ns/op	    2640 B/op	      10 allocs/op
BenchmarkJSONUnmarshal-8           	  500000	      3788 ns/op	    2640 B/op	      10 allocs/op
BenchmarkJSONUnmarshal-8           	  500000	      3330 ns/op	    2640 B/op	      10 allocs/op
BenchmarkJSONUnmarshal-8           	  500000	      3092 ns/op	    2640 B/op	      10 allocs/op
BenchmarkJSONUnmarshal-8           	  500000	      3422 ns/op	    2640 B/op	      10 allocs/op
BenchmarkJSONUnmarshal-8           	  500000	      3224 ns/op	    2640 B/op	      10 allocs/op
BenchmarkJSONUnmarshalParallel-8   	 1000000	      1329 ns/op	    2640 B/op	      10 allocs/op
BenchmarkJSONUnmarshalParallel-8   	 1000000	      1385 ns/op	    2640 B/op	      10 allocs/op
BenchmarkJSONUnmarshalParallel-8   	 1000000	      1401 ns/op	    2640 B/op	      10 allocs/op
BenchmarkJSONUnmarshalParallel-8   	 1000000	      1487 ns/op	    2640 B/op	      10 allocs/op
BenchmarkJSONUnmarshalParallel-8   	 1000000	      1414 ns/op	    2640 B/op	      10 allocs/op
BenchmarkJSONUnmarshalParallel-8   	 1000000	      1463 ns/op	    2640 B/op	      10 allocs/op
BenchmarkJSONUnmarshalParallel-8   	 1000000	      1472 ns/op	    2640 B/op	      10 allocs/op
BenchmarkJSONUnmarshalParallel-8   	 1000000	      1850 ns/op	    2640 B/op	      10 allocs/op
BenchmarkJSONUnmarshalParallel-8   	 1000000	      1789 ns/op	    2640 B/op	      10 allocs/op
BenchmarkJSONUnmarshalParallel-8   	 1000000	      1638 ns/op	    2640 B/op	      10 allocs/op
PASS
ok  	github.com/kpango/json-bench	32.948s
go tool pprof --svg pprof/test.bin pprof/cpu-unmarshal.out > cpu-unmarshal.svg
go tool pprof --svg pprof/test.bin pprof/mem-unmarshal.out > mem-unmarshal.svg
go test -count=10 -run=NONE -bench=BenchmarkJSONDecode -benchmem -o pprof/test.bin -cpuprofile pprof/cpu-decode.out -memprofile pprof/mem-decode.out
goos: darwin
goarch: amd64
pkg: github.com/kpango/json-bench
BenchmarkJSONDecode-8           	  500000	      2630 ns/op	    1192 B/op	       9 allocs/op
BenchmarkJSONDecode-8           	  500000	      2609 ns/op	    1192 B/op	       9 allocs/op
BenchmarkJSONDecode-8           	  500000	      2592 ns/op	    1192 B/op	       9 allocs/op
BenchmarkJSONDecode-8           	  500000	      2568 ns/op	    1192 B/op	       9 allocs/op
BenchmarkJSONDecode-8           	  500000	      2591 ns/op	    1192 B/op	       9 allocs/op
BenchmarkJSONDecode-8           	  500000	      2589 ns/op	    1192 B/op	       9 allocs/op
BenchmarkJSONDecode-8           	  500000	      2595 ns/op	    1192 B/op	       9 allocs/op
BenchmarkJSONDecode-8           	  500000	      2840 ns/op	    1192 B/op	       9 allocs/op
BenchmarkJSONDecode-8           	  500000	      2641 ns/op	    1192 B/op	       9 allocs/op
BenchmarkJSONDecode-8           	  500000	      2548 ns/op	    1192 B/op	       9 allocs/op
BenchmarkJSONDecodeParallel-8   	 2000000	       881 ns/op	    1192 B/op	       9 allocs/op
BenchmarkJSONDecodeParallel-8   	 2000000	       920 ns/op	    1192 B/op	       9 allocs/op
BenchmarkJSONDecodeParallel-8   	 2000000	       963 ns/op	    1192 B/op	       9 allocs/op
BenchmarkJSONDecodeParallel-8   	 2000000	       954 ns/op	    1192 B/op	       9 allocs/op
BenchmarkJSONDecodeParallel-8   	 1000000	      1022 ns/op	    1192 B/op	       9 allocs/op
BenchmarkJSONDecodeParallel-8   	 2000000	      1001 ns/op	    1192 B/op	       9 allocs/op
BenchmarkJSONDecodeParallel-8   	 2000000	       974 ns/op	    1192 B/op	       9 allocs/op
BenchmarkJSONDecodeParallel-8   	 1000000	      1019 ns/op	    1192 B/op	       9 allocs/op
BenchmarkJSONDecodeParallel-8   	 2000000	       977 ns/op	    1192 B/op	       9 allocs/op
BenchmarkJSONDecodeParallel-8   	 2000000	       974 ns/op	    1192 B/op	       9 allocs/op
PASS
ok  	github.com/kpango/json-bench	38.557s
go tool pprof --svg pprof/test.bin pprof/mem-decode.out > mem-decode.svg
go tool pprof --svg pprof/test.bin pprof/cpu-decode.out > cpu-decode.svg
mv ./*.svg bench/
