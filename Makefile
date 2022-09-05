install:
	@GO111MODULE=on go mod tidy
	@GO111MODULE=on go mod vendor
test:
	@go test -v -race ./...
