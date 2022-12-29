lint:
	go vet ./...

build_dir:
	mkdir -p build

test: lint build_dir
	go test -v ./... -covermode=count -coverprofile=coverage.out && go tool cover -func=coverage.out -o=build/coverage.out

build: test
	go build -o build -v ./...