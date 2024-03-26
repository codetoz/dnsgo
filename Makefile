install:
	go get
run:
	go run ./main.go
build:
	make clean && go build -o ./bin/dnsgo
clean:
	rm -rf bin/
