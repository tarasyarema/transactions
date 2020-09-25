run: build
	./transactions.exe

build:
	go build -race -v

test:
	go test -v

bench:
	go test -bench=.
