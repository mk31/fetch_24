package main

import (
	"testing"
)

func TestIsTrimmedLengthMultipleOf3(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"MultipleOf3", "abc", true},
		{"NotMultipleOf3", "ab", false},
		{"OnlySpaces", "   ", false},
		{"SpacesWithMultipleOf3", "  abc  ", true},
		{"SpacesWithNotMultipleOf3", "  ab  ", false},
		{"EmptyString", "", false},
	}

	for _, test := range tests {

		result := IsTrimmedLengthMultipleOf3(test.input)

		if result != test.expected {
			t.Errorf("%v IsTrimmedLengthMultipleOf3(%v) = %v; want %v", test.name, test.input, result, test.expected)
		}
	}
}
