ARG DOCS_MANAGER_BUILD_VERSION=stable

FROM onosproject/golang-build:$DOCS_MANAGER_BUILD_VERSION AS build
ENV GO111MODULE=on
COPY . /go/src/github.com/multi-repo-docs-versioning/docs-manager
RUN cd /go/src/github.com/multi-repo-docs-versioning/docs-manager && GOFLAGS=-mod=vendor make build-docs-manager

FROM alpine:3.9
RUN apk add libc6-compat
RUN apk add --no-cache git
RUN apk add --no-cache bash
RUN apk add --no-cache openssh


USER root

COPY --from=build /go/src/github.com/multi-repo-docs-versioning/docs-manager/build/_output/docs-manager /usr/local/bin/docs-manager

