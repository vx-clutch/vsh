all: build

build:
	@go build -o ./go-out/bin/ ./...

debug:
	@go build -gcflags="all=-N -l" -o ./go-out/bin/ ./...

clean:
	@-rm ./go-out/bin/*
	@go clean
