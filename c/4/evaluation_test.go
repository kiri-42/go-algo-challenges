package main

import (
	"testing"
)

// TestFindEscapeCodeByChatGPT は、FindEscapeCode関数のテストを行います。
func TestFindEscapeCodeByChatGPT(t *testing.T) {
	// テストケースを定義
	testCases := []struct {
		name     string
		numbers  []int
		expected int
		err      bool
	}{
		{"NormalCase1", []int{1, 2, 3, 2, 3, 3}, 3, false},
		{"NormalCase2", []int{4, 4, 5, 5, 6}, 5, false},
		{"ErrorCaseInvalidNumber", []int{10, 2, 3}, 0, true},
		{"ErrorCaseEmptySlice", []int{}, 0, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			code, err := FindEscapeCode(tc.numbers)

			// エラーの発生を検証
			if tc.err && err == nil {
				t.Errorf("Expected an error for %v, but got none", tc.name)
			}
			if !tc.err && err != nil {
				t.Errorf("Did not expect an error for %v, but got: %v", tc.name, err)
			}

			// 戻り値を検証
			if code != tc.expected {
				t.Errorf("Expected %v for %v, but got %v", tc.expected, tc.name, code)
			}
		})
	}
}
