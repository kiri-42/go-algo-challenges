package main

import (
	"testing"
)

// SumNumbersByChatGPT は SumNumbers 関数をテストします。
func TestSumNumbersByChatGPT(t *testing.T) {
	tests := []struct {
		name          string
		numbers       string
		expectedSum   int
		expectedError bool
	}{
		{"NormalCase", "1,2,3", 6, false},
		{"NegativeNumbers", "-1,-2,-3", -6, false},
		{"InvalidNumbers", "a,2,3", 0, true},
		{"EmptyString", "", 0, false},
		{"SingleNumber", "5", 5, false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			sum, err := SumNumbers(test.numbers)
			if (err != nil) != test.expectedError {
				t.Errorf("Test %s failed: expected error status %v, got %v", test.name, test.expectedError, err != nil)
			}
			if sum != test.expectedSum {
				t.Errorf("Test %s failed: expected sum %d, got %d", test.name, test.expectedSum, sum)
			}
		})
	}
}
