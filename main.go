package main

import (
	"gitea.com/hkrd/protoc-gen-gotag/module"
	plugin_go "github.com/golang/protobuf/protoc-gen-go/plugin"
	pgs "github.com/lyft/protoc-gen-star"
	pgsgo "github.com/lyft/protoc-gen-star/lang/go"
	"google.golang.org/protobuf/proto"
)

func main() {
	pgs.Init(
		pgs.DebugEnv("GOTAG_DEBUG"),
		pgs.SupportedFeatures(proto.Uint64(uint64(plugin_go.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL))),
	).RegisterModule(module.New()).
		RegisterPostProcessor(pgsgo.GoFmt()).
		Render()
}
