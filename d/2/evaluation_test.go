package main

import (
	"testing"
)

// TestSumRangeByChatGPT はsumRange関数のテストを行います。
func TestSumRangeByChatGPT(t *testing.T) {
	tests := []struct {
		name    string
		numbers []int
		start   int
		end     int
		want    int
		wantErr bool
	}{
		{"正常なケース1", []int{1, 2, 3, 4, 5}, 1, 3, 9, false},
		{"正常なケース2", []int{10, 20, 30}, 0, 2, 60, false},
		{"エラーケース：範囲逆", []int{1, 2, 3}, 2, 1, 0, true},
		{"エラーケース：範囲外（開始）", []int{1, 2, 3, 4, 5}, -1, 4, 0, true},
		{"エラーケース：範囲外（終了）", []int{1, 2, 3, 4, 5}, 1, 5, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SumRange(tt.numbers, tt.start, tt.end)
			if (err != nil) != tt.wantErr {
				t.Errorf("sumRange() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("sumRange() = %v, want %v", got, tt.want)
			}
		})
	}
}
