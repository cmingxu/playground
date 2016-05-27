
all: run

GO=go

run: build
	@./bin/run

build:
	@mkdir -p bin
	@$(GO) build -o bin/run main.go

clean:
	rm -rf bin

