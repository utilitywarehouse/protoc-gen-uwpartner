package main

import (
	pgs "github.com/lyft/protoc-gen-star"
	pgsgo "github.com/lyft/protoc-gen-star/lang/go"
	"github.com/utilitywarehouse/protoc-gen-uwpartner/service"
)

func main() {
	pgs.Init(
		pgs.DebugEnv("DEBUG"),
	).RegisterModule(
		service.Module(),
	).RegisterPostProcessor(
		pgsgo.GoFmt(),
	).Render()
}
