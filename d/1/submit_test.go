package main

import (
	"errors"
	"testing"
)

func TestSumNumbers(t *testing.T) {
	type input struct {
		numbers string
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
			name:   "正常時: 有効な数字のみで構成された場合、合計を返すこと",
			input:  input{"2,8,56"},
			output: output{66, nil},
		},
		{
			name:   "正常時: 負の数を含む有効な数字のみで構成された場合、合計を返すこと",
			input:  input{"2,-8,56"},
			output: output{50, nil},
		},
		{
			name:   "正常時: 空文字列の場合、0を返すこと",
			input:  input{""},
			output: output{0, nil},
		},
		{
			name:   "異常時: 無効な文字が含まれていた場合、エラーを返すこと",
			input:  input{"2,8,a"},
			output: output{0, errors.New("invalid number: a")},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result, err := SumNumbers(c.input.numbers)
			if result != c.output.sum {
				t.Errorf("SumNumbers(%v) result %v expected %v", c.input.numbers, result, c.output.sum)
			}
			if c.output.err == nil {
				if err != nil {
					t.Errorf("SumNumbers(%v) result(error) %v expected %v", c.input.numbers, err, c.output.err)
				}
			} else {
				if err != nil && err.Error() != c.output.err.Error() {
					t.Errorf("SumNumbers(%v) result(error) %v expected %v", c.input.numbers, err, c.output.err)
				}
			}
		})
	}
}
