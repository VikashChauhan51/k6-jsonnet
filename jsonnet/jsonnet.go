package jsonnet

import (
	gj "github.com/google/go-jsonnet"
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
