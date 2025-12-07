default:
    @just --list

encode:
    go run ./cmd/encode

decode:
    go run ./cmd/decode

encode_binary_name := `awk -F'"' '/EncodeBinaryName.*=/ {print $2}' internal/config/constants.go`
decode_binary_name := `awk -F'"' '/DecodeBinaryName.*=/ {print $2}' internal/config/constants.go`

build:
    go build -o bin/{{encode_binary_name}} ./cmd/encode
    go build -o bin/{{decode_binary_name}} ./cmd/decode

test: vet
    go test -v -coverpkg=./internal -coverprofile=coverage.out ./...
    go tool cover -html=coverage.out -o coverage.html

integration-test: vet            
    go test -v -tags=integration -coverpkg=./internal -coverprofile=coverage.integration.out ./tests
    go tool cover -html=coverage.integration.out -o coverage.integration.html

benchmark:                                        
    go test -bench=. ./... | tee benchmark.out

format:
    go fmt ./...

vet:
    go vet ./...

check: format test integration-test    

clean:
    rm -rf bin/
    rm -f coverage*.out coverage*.html benchmark*.out
    go clean
