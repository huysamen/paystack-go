// Package optional provides utility functions for implementing functional options pattern
// across the Paystack Go SDK.
package optional

import "time"

// Helper functions for creating pointers to primitive types
// These are useful for optional fields in request structs

// String returns a pointer to the given string value
func String(s string) *string {
	return &s
}

// Int returns a pointer to the given int value
func Int(i int) *int {
	return &i
}

// Uint64 returns a pointer to the given uint64 value
func Uint64(u uint64) *uint64 {
	return &u
}

// Bool returns a pointer to the given bool value
func Bool(b bool) *bool {
	return &b
}

// Float64 returns a pointer to the given float64 value
func Float64(f float64) *float64 {
	return &f
}

// Time returns a pointer to the given time.Time value
func Time(t time.Time) *time.Time {
	return &t
}
