package jsonnet

import (
	"encoding/json"
	"testing"
)

func TestEvaluateJsonnet(t *testing.T) {
	// Create a new Extension
	ext := New()

	// Define a simple Jsonnet template
	jsonnetTemplate := `{ "hello": "world" }`

	// Evaluate the Jsonnet template
	jsonStr, err := ext.EvaluateJsonnet(jsonnetTemplate)

	// Check for errors
	if err != nil {
		t.Errorf("EvaluateJsonnet returned an error: %v", err)
	}

	// Convert the JSON object to a JSON string
	jsonBytes, err := json.Marshal(jsonStr)
	if err != nil {
		t.Errorf("Error marshalling JSON: %v", err)
	}

	// Check the result
	expected := "{\n   \"hello\": \"world\"\n}\n"
	if jsonStr != expected {
		t.Errorf("EvaluateJsonnet returned %v, want %v", string(jsonBytes), expected)
	}
}
