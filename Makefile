build:
	go build -o bin/go-mail-local cmd/main.go

run:
	go run cmd/main.go

clean:
	rm -rf bin/go-mail-local
	rm -rf config
