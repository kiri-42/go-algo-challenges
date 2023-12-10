package main

import (
	"errors"
	"testing"
)

func TestFindMaxValue(t *testing.T) {
	type input struct {
		numbers []int
	}

	type output struct {
		maxValue int
		err      error
	}

	cases := []struct {
		name   string
		input  input
		output output
	}{
		{
			name:   "正常時: 有効な数列の場合、最大値を返すこと",
			input:  input{[]int{4, 0, -54, 307}},
			output: output{307, nil},
		},
		{
			name:   "正常時: 数列が全て負の数の場合、最大値を返すこと",
			input:  input{[]int{-4, -1, -54, -307}},
			output: output{-1, nil},
		},
		{
			name:   "異常時: 有効範囲より小さい値が含まれていた場合、エラーを返すこと",
			input:  input{[]int{4, 0, -54, 307, -1001}},
			output: output{0, errors.New("invalid value in slice")},
		},
		{
			name:   "異常時: 有効範囲より大きい値が含まれていた場合、エラーを返すこと",
			input:  input{[]int{4, 0, -54, 307, 1001}},
			output: output{0, errors.New("invalid value in slice")},
		},
		{
			name:   "異常時: 空のスライスの場合、エラーを返すこと",
			input:  input{[]int{}},
			output: output{0, errors.New("empty slice")},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result, err := FindMaxValue(c.input.numbers)
			if result != c.output.maxValue {
				t.Errorf("FindMaxValue(%v) result %v expected %v", c.input.numbers, result, c.output.maxValue)
			}
			if c.output.err == nil {
				if err != nil {
					t.Errorf("FindMaxValue(%v) result(error) %v expected %v", c.input.numbers, err, c.output.err)
				}
			} else {
				if err.Error() != c.output.err.Error() {
					t.Errorf("FindMaxValue(%v) result(error) %v expected %v", c.input.numbers, err, c.output.err)
				}
			}
		})
	}
}
