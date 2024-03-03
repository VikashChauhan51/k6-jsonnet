package main

import (
	"github.com/VikashChauhan51/k6-jsonnet/jsonnet"
	"go.k6.io/k6/js/modules"
)

func init() {
	modules.Register("k6/x/jsonnet", jsonnet.New())
}
