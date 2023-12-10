package main

import (
	"testing"
)

func TestIsPalindromeByChatGPT(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
		err      bool
	}{
		{"racecar", true, false},
		{"hello", false, false},
		{"あいうえお", false, true},
		{"12321", true, false},
		{"", false, false},
	}

	for _, tc := range testCases {
		result, err := IsPalindrome(tc.input)
		if result != tc.expected {
			t.Errorf("IsPalindrome(%s) = %v, want %v", tc.input, result, tc.expected)
		}
		if (err != nil) != tc.err {
			t.Errorf("IsPalindrome(%s) expected error = %v, got %v", tc.input, tc.err, err)
		}
	}
}
