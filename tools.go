//go:build tools
// +build tools

// nolint
// https://github.com/golang/go/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module

package tools

import (
	_ "github.com/golang/protobuf/protoc-gen-go"
	_ "github.com/ory/go-acc" // accurate code coverage
	// use this for go proto v2
	// _ "google.golang.org/grpc/cmd/protoc-gen-go-grpc"
	// _ "google.golang.org/protobuf/cmd/protoc-gen-go"
)
