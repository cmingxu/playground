
all: build

GO=go

run:
	@./bin/run

build:
	@mkdir -p bin
	@$(GO) build -o bin/run main.go

clean:
	rm -rf bin

