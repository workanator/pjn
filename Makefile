test_all: test benchmark

test:
	@go test -v;

benchmark:
	@go test benchmarks/*.go -v -bench=. -benchmem;
