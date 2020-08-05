install:
	@GO111MODULE=on go mod vendor
test:
	@go test -v -race ./...