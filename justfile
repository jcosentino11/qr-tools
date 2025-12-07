default:
    @just --list

qr-decode:
    go run ./cmd/decode

qr-encode:
    go run ./cmd/encode

build:
    go build -o bin/qr-decode ./cmd/decode
    go build -o bin/qr-encode ./cmd/encode

test:
    go test -v -coverpkg=./internal -coverprofile=coverage.out ./...
    go tool cover -html=coverage.out -o coverage.html

integration-test:                
    go test -v -tags=integration -coverpkg=./internal -coverprofile=coverage.integration.out ./tests
    go tool cover -html=coverage.integration.out -o coverage.integration.html

benchmark:                                        
    go test -bench=. ./... | tee benchmark.out

fmt:
    go fmt ./...

vet:
    go vet ./...

check: fmt vet test integration-test

clean:
    rm -rf bin/
    rm -f coverage*.out coverage*.html benchmark*.out
    go clean
