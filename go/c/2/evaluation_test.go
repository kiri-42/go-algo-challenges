package main

import (
	"testing"
)

func TestFindMaxValueByChatGPT(t *testing.T) {
	testCases := []struct {
		name    string
		numbers []int
		want    int
		wantErr bool
	}{
		{"PositiveNumbers", []int{3, 1, 4, 1, 5, 9}, 9, false},
		{"NegativeNumbers", []int{-1, -3, -10, -2}, -1, false},
		{"EmptySlice", []int{}, 0, true},
		{"InvalidValue", []int{1001, 3, 4}, 0, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := FindMaxValue(tc.numbers)
			if (err != nil) != tc.wantErr {
				t.Errorf("FindMaxValue() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
			if got != tc.want {
				t.Errorf("FindMaxValue() = %v, want %v", got, tc.want)
			}
		})
	}
}
