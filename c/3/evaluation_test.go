package main

import (
	"testing"
)

// TestFindMinAndIndexByChatGPT is the test function for FindMinAndIndex
func TestFindMinAndIndexByChatGPT(t *testing.T) {
	testCases := []struct {
		name          string
		nums          []int
		expectedMin   int
		expectedIndex int
		expectedErr   string
	}{
		{"NormalCase", []int{3, 1, 4, 1, 5, 9}, 1, 1, ""},
		{"AnotherNormalCase", []int{10, 20, 30}, 10, 0, ""},
		{"EmptySlice", []int{}, 0, 0, "invalid input: slice cannot be empty"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			min, index, err := FindMinAndIndex(tc.nums)

			// Check for error message
			if err != nil && err.Error() != tc.expectedErr {
				t.Errorf("FindMinAndIndex(%v) returned unexpected error: got %v, want %v", tc.nums, err, tc.expectedErr)
			} else if err == nil && tc.expectedErr != "" {
				t.Errorf("FindMinAndIndex(%v) expected error %v, got no error", tc.nums, tc.expectedErr)
			}

			// Check for value correctness
			if err == nil && (min != tc.expectedMin || index != tc.expectedIndex) {
				t.Errorf("FindMinAndIndex(%v) = %v, %v; want %v, %v", tc.nums, min, index, tc.expectedMin, tc.expectedIndex)
			}
		})
	}
}
