run: build
	@./bin/aoc2022

build:
	@go build -o bin/aoc2022

clean:
	@rm -rf bin