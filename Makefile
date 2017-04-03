PACKAGES = $(shell go list ./...)

fmt:
	go fmt $(PACKAGES)

test:
	go get gopkg.in/check.v1
	go get github.com/spf13/viper
	go test ./...
