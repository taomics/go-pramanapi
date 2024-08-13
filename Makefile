PROTO_GO_PKG := github.com/taomics/pramanapi
PROTO_OUT := .
API_URL := https://github.com/taomics/pramanapi/archive/refs/heads/main.zip
GO_TEST :=

.PHONY: all

all: protoc

.PHONY: tools

GOBIN := $(PWD)/tools
PATH := $(GOBIN):$(PATH)

export GOBIN
export PATH

tools:
	$(MAKE) -C tools

.PHONY: protoc

PROTOC := protoc \
					-I=tmp/pramanapi \
					-I=tmp/pramanapi/accounts \
					-I=tmp/pramanapi/lifestylejournal \
					-I=tmp/pramanapi/recordlog \
					--go_out=$(PROTO_OUT) --go_opt=paths=source_relative  \
					--go-grpc_out=$(PROTO_OUT) --go-grpc_opt=paths=source_relative

protoc: tmp/pramanapi tools
	$(PROTOC) tmp/pramanapi/*.proto
	$(PROTOC) tmp/pramanapi/accounts/*.proto
	$(PROTOC) tmp/pramanapi/lifestylejournal/*.proto
	$(PROTOC) tmp/pramanapi/recordlog/*.proto
	go mod init $(PROTO_GO_PKG); go mod tidy

tmp/pramanapi: tmp/pramanapi.zip
	cd tmp; unzip pramanapi.zip
	cd tmp; mv pramanapi-main pramanapi
	touch tmp/pramanapi

tmp/pramanapi.zip:
	- mkdir tmp
	curl -L -o tmp/pramanapi.zip $(API_URL)

.PHONY: clean
clean:
	- rm -rf tmp
	- find . -name '*.pb.go' -delete
