PACKAGES = $(shell go list ./...)

fmt:
	go fmt $(PACKAGES)

test:
	go get gopkg.in/check.v1
	go test ./...
