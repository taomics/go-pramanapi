VERSION := 0.20250814.0
API_URL := https://github.com/taomics/pramanapi/archive/refs/tags/v$(VERSION).zip
PRAMANAPI_DIR := ./tmp/pramanapi

.PHONY: all

all: protoc

.PHONY: protoc

PROTOC := docker run --rm -v $(PWD):/workspace/go-pramanapi -v $(PRAMANAPI_DIR):/workspace/pramanapi -w /workspace --platform=linux/amd64 protoc:latest \
	-I=pramanapi \
	-I=pramanapi/accounts \
	-I=pramanapi/lifestylejournal \
	-I=pramanapi/recordlog \
	-I=pramanapi/externaldata \
	--go_out=go-pramanapi --go_opt=paths=source_relative  \
	--go-grpc_out=go-pramanapi --go-grpc_opt=paths=source_relative \
	--connect-go_out=go-pramanapi --connect-go_opt=paths=source_relative

protoc: .protoc-image-version $(PRAMANAPI_DIR)
	$(PROTOC) 'pramanapi/*.proto'
	$(PROTOC) 'pramanapi/accounts/*.proto'
	$(PROTOC) 'pramanapi/lifestylejournal/*.proto'
	$(PROTOC) 'pramanapi/recordlog/*.proto'
	$(PROTOC) 'pramanapi/externaldata/*.proto'
	$(PROTOC) 'pramanapi/healthcheck/*.proto'
	$(PROTOC) 'pramanapi/healthfeedback/*.proto'
	go mod init github.com/taomics/pramanapi; go mod tidy

tmp/pramanapi: tmp/pramanapi.zip
	cd tmp; unzip pramanapi.zip
	cd tmp; mv pramanapi-$(VERSION) pramanapi
	touch tmp/pramanapi

tmp/pramanapi.zip:
	- mkdir tmp
	curl -L -o tmp/pramanapi.zip $(API_URL)

.PHONY: image
image: .protoc-image-version

.protoc-image-version: Dockerfile docker-entrypoint.sh .dockerignore
	docker build --platform=linux/amd64 -t protoc:latest .
	docker images -q protoc:latest > .protoc-image-version

.PHONY: clean
clean:
	- rm -rf tmp
	- find . -name '*.pb.go' -delete
