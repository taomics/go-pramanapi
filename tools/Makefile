.PHONY: tools

GOBIN := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))
COMMANDS := protoc-gen-go protoc-gen-go-grpc

export GOBIN

tools: $(COMMANDS)
	@./protoc-gen-go --version
	@./protoc-gen-go-grpc --version

$(COMMANDS): go.mod
	go generate

.PHONY: clean

clean:
	- rm protoc-gen-go
	- rm protoc-gen-go-grpc
