VERSION ?= $(shell git describe --tags --abbrev=0)
LDFLAGS = "-linkmode internal -X ${REPO}/app.VERSION=${VERSION}"


.PHONY: build_server
build_server: BUILDFLAGS=${LDFLAGS}
build_server:
	go build -o bin/server \
 		-ldflags ${BUILDFLAGS} \
		./cmd/server/main.go

.PHONY: build_client
build_client: BUILDFLAGS=${LDFLAGS}
build_client:
	go build -o bin/client \
		-ldflags ${BUILDFLAGS} \
		./cmd/client/main.go

.PHONY: build
build: build_server build_client
