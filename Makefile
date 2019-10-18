export CGO_ENABLED=0
export GO111MODULE=on

DOCS_MANAGER_BUILD_VERSION := stable
DOCS_MANAGER_TEST_VERSION := latest

all: build

build: deps build-docs-manager linters

build-docs-manager: 
	go build -o build/_output/docs-manager ./cmd/docs-manager

images:
	@go mod vendor
	docker build . -f build/docs-manager/Dockerfile \
		--build-arg DOCS_MANAGER_BUILD_VERSION=${DOCS_MANAGER_BUILD_VERSION} \
		-t docs-manager:${DOCS_MANAGER_TEST_VERSION}
	@rm -rf vendor

linters: # @HELP examines Go source code and reports coding problems
	golangci-lint run

deps: # @HELP ensure that the required dependencies are in place
	#go get -u -v gopkg.in/src-d/go-git.v4/...
	go build -v ./...
	bash -c "diff -u <(echo -n) <(git diff go.mod)"
	bash -c "diff -u <(echo -n) <(git diff go.sum)"
