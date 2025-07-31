package utils

import (
	"fmt"
	"regexp"
)

// ValidationError represents a validation error with detailed context
type ValidationError struct {
	Field   string
	Value   any
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation error for field '%s': %s", e.Field, e.Message)
}

// NewValidationError creates a new validation error
func NewValidationError(field string, value any, message string) *ValidationError {
	return &ValidationError{
		Field:   field,
		Value:   value,
		Message: message,
	}
}

// ValidateEmail validates an email address format
func ValidateEmail(email string) error {
	if email == "" {
		return NewValidationError("email", email, "email is required")
	}

	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(email) {
		return NewValidationError("email", email, "invalid email format")
	}

	return nil
}

// ValidateAmount validates a monetary amount
func ValidateAmount(amount int) error {
	if amount <= 0 {
		return NewValidationError("amount", amount, "amount must be greater than 0")
	}

	return nil
}

// ValidateRequired validates that a string field is not empty
func ValidateRequired(field, value string) error {
	if value == "" {
		return NewValidationError(field, value, fmt.Sprintf("%s is required", field))
	}

	return nil
}

// CombineValidationErrors combines multiple validation errors into one
func CombineValidationErrors(errors ...error) error {
	var validationErrs []*ValidationError

	for _, err := range errors {
		if err != nil {
			if valErr, ok := err.(*ValidationError); ok {
				validationErrs = append(validationErrs, valErr)
			} else {
				// Convert regular errors to validation errors
				validationErrs = append(validationErrs, &ValidationError{
					Field:   "unknown",
					Message: err.Error(),
				})
			}
		}
	}

	if len(validationErrs) == 0 {
		return nil
	}

	if len(validationErrs) == 1 {
		return validationErrs[0]
	}

	// Combine multiple errors
	var messages []string
	for _, err := range validationErrs {
		messages = append(messages, err.Error())
	}

	return fmt.Errorf("multiple validation errors: %v", messages)
}
