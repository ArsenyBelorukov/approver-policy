//go:build tools
// +build tools

// This file exists to force 'go mod' to fetch tool dependencies
// See: https://github.com/golang/go/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module

package bin

import (
	_ "github.com/norwoodj/helm-docs/cmd/helm-docs"
	_ "github.com/onsi/ginkgo/v2/ginkgo"
	_ "github.com/princjef/gomarkdoc/cmd/gomarkdoc"
	_ "google.golang.org/protobuf/cmd/protoc-gen-go"
	_ "k8s.io/code-generator/cmd/deepcopy-gen"
	_ "sigs.k8s.io/controller-tools/cmd/controller-gen"
	_ "sigs.k8s.io/kind"
	_ "sigs.k8s.io/kustomize/kustomize/v5"
)
