install:
	go get ./cmd/
run:
	go run ./main.go
build:
	make clean && go build -o ./bin/dnsgo
clean:
	rm -rf bin/
