package path

import (
	"encoding/json"
	"reflect"
	"testing"
)

// TestError tests generic errors returned from Path's API
func TestError(t *testing.T) {
	const jsonError = `{"detail":"Not authenticated"}`

	expected := Error{Detail: "Not authenticated"}

	var got Error
	err := json.Unmarshal([]byte(jsonError), &got)
	if err != nil {
		t.Errorf("Error unmarshalling JSON into struct: %s\n", err.Error())
	}

	if got != expected {
		t.Errorf("Expected %+v, got %+v\n", expected, got)
	}
}

// TestValidationError ensures that the library properly handles ValidationError responses
func TestValidationError(t *testing.T) {
	const jsonError = `{"detail":[{"loc":["body","password"],"msg":"field required","type":"value_error.missing"}]}`

	expected := ValidationError{
		Detail: []ValidationErrorItem{
			{
				Loc:  []string{"body", "password"},
				Msg:  "field required",
				Type: "value_error.missing",
			},
		},
	}

	var got ValidationError
	err := json.Unmarshal([]byte(jsonError), &got)
	if err != nil {
		t.Errorf("Error unmarshalling JSON into struct: %s\n", err.Error())
	}

	// DeepEqual is used to recursively compare all elements
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Expected %+v, got %+v\n", expected, got)
	}
}

// TestHTTPValidationError makes sure that HTTPValidationErrors are marshaled correctly
func TestHTTPValidationError(t *testing.T) {
	const jsonError = `{"detail":[{"loc":["body","password"],"msg":"field required","type":"value_error.missing"}]}`

	expected := HTTPValidationError{
		Detail: []ValidationErrorItem{
			{
				Loc:  []string{"body", "password"},
				Msg:  "field required",
				Type: "value_error.missing",
			},
		},
	}

	var got HTTPValidationError
	err := json.Unmarshal([]byte(jsonError), &got)
	if err != nil {
		t.Errorf("Error unmarshalling JSON into struct: %s\n", err.Error())
	}

	// DeepEqual is used to recursively compare all elements
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Expected %+v, got %+v\n", expected, got)
	}
}
