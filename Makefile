install:
	@GO111MODULE=on go mod vendor
start:
	@go run *.go
build:
	@GOOS=linux GOARCH=amd64 go build -v ./*.go
clean:
	@rm -rf main
test:
	@go test -v -race ./...