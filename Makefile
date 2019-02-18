test_all: test benchmark

test:
	@go test tests/*.go -v;

benchmark:
	@go test benchmarks/*.go -v -bench=. -benchmem;
