package main

import (
	"testing"
	"errors"
)

func TestIsPalindrome(t *testing.T) {
	type input struct {
		s string
	}

	type output struct {
		isPal bool
		err   error
	}

	cases := []struct {
		name   string
		input  input
		output output
	}{
		{
			name: "正常時: 回文",
			input: input{"racecar"},
			output: output{true, nil},
		},
		{
			name: "正常時: 回文でない",
			input: input{"hello"},
			output: output{false, nil},
		},
		{
			name: "異常時: 空文字",
			input: input{""},
			output: output{false, nil},
		},
		{
			name: "異常時: スペースが含まれる",
			input: input{"hello world"},
			output: output{false, errors.New("英数字以外が含まれています")},
		},
		{
			name: "異常時: 日本語が含まれる",
			input: input{"あいうえお"},
			output: output{false, errors.New("英数字以外が含まれています")},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result, err := IsPalindrome(c.input.s)
			if result != c.output.isPal {
				t.Errorf("IsPalindrome(%v) result %v expected %v", c.input.s, result, c.output.isPal)
			}
			if c.output.err == nil{
				if err != nil {
					t.Errorf("IsPalindrome(%v) result %v expected %v", c.input.s, err, c.output.err)
				}
			} else {
				if err == nil {
					t.Errorf("IsPalindrome(%v) result %v expected %v", c.input.s, err, c.output.err)
				}
			}
		})
	}
}
