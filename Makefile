all: build

build: deps build-docs-manager linters

build-docs-manager: 
	go build -o build/_output/docs-manager ./cmd/docs-manager


linters: # @HELP examines Go source code and reports coding problems
	golangci-lint run

deps: # @HELP ensure that the required dependencies are in place
	go get -u -v gopkg.in/src-d/go-git.v4/...
	go build -v ./...
	bash -c "diff -u <(echo -n) <(git diff go.mod)"
	bash -c "diff -u <(echo -n) <(git diff go.sum)"
