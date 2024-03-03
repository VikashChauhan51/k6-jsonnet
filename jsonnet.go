package jsonnet

import (
	gj "github.com/google/go-jsonnet"
	"go.k6.io/k6/js/modules"
)

type Extension struct{}

func New() *Extension {
	return &Extension{}
}

func (e *Extension) EvaluateJsonnet(jsonnetTemplate string) (string, error) {
	vm := gj.MakeVM()
	jsonStr, err := vm.EvaluateSnippet("snippet", jsonnetTemplate)
	if err != nil {
		return "", err
	}
	return jsonStr, nil
}

// Register the extensions on module initialization.
func init() {
	modules.Register("k6/x/jsonnet", New())
}
