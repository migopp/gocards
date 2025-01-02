up:
	go run ./cmd/gocards/main.go

test:
	go test ./... -v

cleandb:
	find . -type f -name "*.db" -exec rm -f {} +

clean:

.PHONY: all test cleandb clean
