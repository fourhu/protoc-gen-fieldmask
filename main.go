package main

import (
	"github.com/fourhu/protoc-gen-fieldmask/internal/module"

	pgs "github.com/lyft/protoc-gen-star"
	pgsgo "github.com/lyft/protoc-gen-star/lang/go"
	"google.golang.org/protobuf/types/pluginpb"
)

func main() {
	optional := uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)

	pgs.Init(pgs.DebugEnv("DEBUG_PGFM"), pgs.SupportedFeatures(&optional)).
		RegisterModule(module.FieldMask(tplRegistryFactory)).
		RegisterPostProcessor(
			pgsgo.GoFmt(),
		).
		Render()
}
