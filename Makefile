GO_VERSION:=$(shell go version)

.PHONY: bench profile clean

all: install

test:
	go test -v -race ./...

bench:
	go test -count=5 -run=NONE -bench . -benchmem

profile: clean
	mkdir bench
	mkdir pprof
	go test -count=10 -run=NONE -bench . -benchmem -o pprof/test.bin -cpuprofile pprof/cpu.out -memprofile pprof/mem.out
	go tool pprof --svg pprof/test.bin pprof/mem.out > mem.svg
	go tool pprof --svg pprof/test.bin pprof/cpu.out > cpu.svg
	go test -count=10 -run=NONE -bench=BenchmarkJSONUnmarshal -benchmem -o pprof/test.bin -cpuprofile pprof/cpu-unmarshal.out -memprofile pprof/mem-unmarshal.out
	go tool pprof --svg pprof/test.bin pprof/cpu-unmarshal.out > cpu-unmarshal.svg
	go tool pprof --svg pprof/test.bin pprof/mem-unmarshal.out > mem-unmarshal.svg
	go test -count=10 -run=NONE -bench=BenchmarkJSONDecode -benchmem -o pprof/test.bin -cpuprofile pprof/cpu-decode.out -memprofile pprof/mem-decode.out
	go tool pprof --svg pprof/test.bin pprof/mem-decode.out > mem-decode.svg
	go tool pprof --svg pprof/test.bin pprof/cpu-decode.out > cpu-decode.svg
	mv ./*.svg bench/

clean:
	rm -rf pprof
	rm -rf bench
