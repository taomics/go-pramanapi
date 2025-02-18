package tools

import (
	_ "connectrpc.com/connect"
	_ "google.golang.org/grpc/cmd/protoc-gen-go-grpc"
	_ "google.golang.org/protobuf"
)

//go:generate go install google.golang.org/protobuf/cmd/protoc-gen-go
//go:generate	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
//go:generate	go install connectrpc.com/connect/cmd/protoc-gen-connect-go
