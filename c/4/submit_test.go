package main

import (
	"errors"
	"testing"
)

func TestFindEscapeCode(t *testing.T) {
	type input struct {
		numbers []int
	}

	type output struct {
		code int
		err  error
	}

	cases := []struct {
		name   string
		input  input
		output output
	}{
		{
			name:   "正常時: 有効な値の場合、最も出現頻度の多い値を返すこと",
			input:  input{[]int{2, 4, 9, 1, 4, 4}},
			output: output{4, nil},
		},
		{
			name:   "正常時: 最も出現頻度の多い値が複数ある場合、最も大きい値を返すこと",
			input:  input{[]int{2, 4, 9, 1, 4, 9}},
			output: output{9, nil},
		},
		{
			name:   "正常時: 値が1つしかない場合、その値を返すこと",
			input:  input{[]int{2}},
			output: output{2, nil},
		},
		{
			name:   "異常時: 無効な値(-1)が含まれていた場合、エラーを返すこと",
			input:  input{[]int{-1, 4}},
			output: output{0, errors.New("invalid number value")},
		},
		{
			name:   "異常時: 無効な値(10)が含まれていた場合、エラーを返すこと",
			input:  input{[]int{10, 4}},
			output: output{0, errors.New("invalid number value")},
		},
		{
			name:   "異常時: スライスが空の場合、エラーを返すこと",
			input:  input{[]int{}},
			output: output{0, errors.New("no numbers provided")},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result, err := FindEscapeCode(c.input.numbers)
			if result != c.output.code {
				t.Errorf("FindEscapeCode(%v) result %v %v expected %v %v", c.input.numbers, result, err, c.output.code, c.output.err)
				return
			}
			if c.output.err == nil {
				if err != nil {
					t.Errorf("FindEscapeCode(%v) result %v %v expected %v %v", c.input.numbers, result, err, c.output.code, c.output.err)
				}
			} else {
				if err != nil && err.Error() != c.output.err.Error() {
					t.Errorf("FindEscapeCode(%v) result %v %v expected %v %v", c.input.numbers, result, err, c.output.code, c.output.err)
				}
			}
		})
	}
}
