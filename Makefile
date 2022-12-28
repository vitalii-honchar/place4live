lint:
	go vet ./...

test: lint
	go test -v ./... -covermode=count -coverprofile=coverage.out && go tool cover -func=coverage.out -o=coverage.out

build: test
	mkdir -p build && go build -o build -v ./...