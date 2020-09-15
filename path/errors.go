package path

// Represents a generic error
type Error struct {
	Detail string `json:"detail"`
}

// HTTPValidationError is an error returned when an expected value is missing or malformed
type HTTPValidationError struct {
	Detail []ValidationErrorItem `json:"detail"`
}

// ValidationError outlines the details of the validation issue that occurred
type ValidationError struct {
	Detail []ValidationErrorItem `json:"detail"`
}

// ValidationError outlines the details of the validation issue that occurred
type ValidationErrorItem struct {
	// Where the ValidationError happened
	Loc  []string `json:"loc"`
	// A message to accommodate the error
	Msg  string   `json:"msg"`
	// The type of ValidationError that occurred
	Type string   `json:"type"`
}