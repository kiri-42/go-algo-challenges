package main

import (
	"errors"
	"testing"
)

func TestFindMinAndIndex(t *testing.T) {
	type input struct {
		nums []int
	}

	type output struct {
		min   int
		index int
		err   error
	}

	cases := []struct {
		name   string
		input  input
		output output
	}{
		{
			name:   "正常時: 有効な値の場合、最小値とその位置を返すこと",
			input:  input{[]int{4, 17, -2, 8, -44}},
			output: output{-44, 4, nil},
		},
		{
			name:   "正常時: 同じ値が複数ある場合、最小値とその位置を返すこと",
			input:  input{[]int{0, 0, 0}},
			output: output{0, 0, nil},
		},
		{
			name:   "異常時: スライスが空の場合、エラーを返すこと",
			input:  input{[]int{}},
			output: output{0, 0, errors.New("invalid input: slice cannot be empty")},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			min, index, err := FindMinAndIndex(c.input.nums)
			if min != c.output.min || index != c.output.index {
				t.Errorf("FindMinAndIndex(%v) result %v %v %v expected %v %v %v", c.input.nums, min, index, err, c.output.min, c.output.index, c.output.err)
				return
			}
			if c.output.err == nil {
				if err != nil {
					t.Errorf("FindMinAndIndex(%v) result %v %v %v expected %v %v %v", c.input.nums, min, index, err, c.output.min, c.output.index, c.output.err)
				}
				return
			} else {
				if err != nil && err.Error() != c.output.err.Error() {
					t.Errorf("FindMinAndIndex(%v) result %v %v %v expected %v %v %v", c.input.nums, min, index, err, c.output.min, c.output.index, c.output.err)
				}
				return
			}
		})
	}
}
