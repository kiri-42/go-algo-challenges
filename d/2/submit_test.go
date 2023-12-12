package main

import (
	"errors"
	"testing"
)

func TestSumRange(t *testing.T) {
	type input struct {
		numbers []int
		start   int
		end     int
	}

	type output struct {
		sum int
		err error
	}

	cases := []struct {
		name   string
		input  input
		output output
	}{
		{
			name:   "正常時: 全ての入力値が有効な場合、範囲内の合計を返すこと",
			input:  input{[]int{2, 1, 7, 11, 4, 6}, 1, 4},
			output: output{23, nil},
		},
		{
			name:   "正常時: startが0の場合、範囲内の合計を返すこと",
			input:  input{[]int{2, 1, 7, 11, 4, 6}, 0, 4},
			output: output{25, nil},
		},
		{
			name:   "正常時: endが要素数-1の場合、範囲内の合計を返すこと",
			input:  input{[]int{2, 1, 7, 11, 4, 6}, 1, 5},
			output: output{29, nil},
		},
		{
			name:   "異常時: startが無効な値の場合、エラーを返すこと",
			input:  input{[]int{2, 1, 7, 11, 4, 6}, -1, 4},
			output: output{0, errors.New("error")},
		},
		{
			name:   "異常時: endが無効な値の場合、エラーを返すこと",
			input:  input{[]int{2, 1, 7, 11, 4, 6}, 1, 6},
			output: output{0, errors.New("error")},
		},
		{
			name:   "異常時: endがstartより小さい場合、エラーを返すこと",
			input:  input{[]int{2, 1, 7, 11, 4, 6}, 1, 0},
			output: output{0, errors.New("error")},
		},
		{
			name:   "異常時: numbersの要素数が無効な場合、エラーを返すこと",
			input:  input{[]int{}, 1, 4},
			output: output{0, errors.New("error")},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result, err := SumRange(c.input.numbers, c.input.start, c.input.end)
			if result != c.output.sum {
				t.Errorf("SumRange(%v, %v, %v) result %v %v expected %v %v", c.input.numbers, c.input.start, c.input.end, result, err, c.output.sum, c.output.err)
			}
			if c.output.err == nil {
				if err != nil {
					t.Errorf("SumRange(%v, %v, %v) result %v %v expected %v %v", c.input.numbers, c.input.start, c.input.end, result, err, c.output.sum, c.output.err)
				}
			} else {
				if err != nil && err.Error() != c.output.err.Error() {
					t.Errorf("SumRange(%v, %v, %v) result %v %v expected %v %v", c.input.numbers, c.input.start, c.input.end, result, err, c.output.sum, c.output.err)
				}
			}
		})
	}
}
