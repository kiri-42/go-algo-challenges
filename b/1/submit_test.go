package main

import (
	"errors"
	"testing"
)

func TestCalculateOptimalPath(t *testing.T) {
	type input struct {
		grid           [][]int
		startX, startY int
		endX, endY     int
	}

	type output struct {
		totalCost int
		err       error
	}

	cases := []struct {
		name   string
		input  input
		output output
	}{
		{
			name: "正常時: 有効な値の場合、最小コストの経路の合計コストを返すこと",
			input: input{
				[][]int{
					{3, 7, 1},
					{6, 3, 4},
					{7, 2, 1},
				}, 0, 0, 2, 1,
			},
			output: output{11, nil},
		},
		{
			name: "正常時: startとendが同じ場合、0を返すこと",
			input: input{
				[][]int{
					{3, 7, 1},
					{6, 3, 4},
					{7, 2, 1},
				}, 0, 0, 0, 0,
			},
			output: output{0, nil},
		},
		{
			name: "異常時: girdの値に0が含まれていた場合、エラーを返すこと",
			input: input{
				[][]int{
					{3, 7, 1},
					{6, 3, 4},
					{7, 2, 0},
				}, 0, 0, 2, 1,
			},
			output: output{0, errors.New("invalid grid value")},
		},
		{
			name: "異常時: startかendの座標がgirdの範囲外の場合、エラーを返すこと",
			input: input{
				[][]int{
					{3, 7, 1},
					{6, 3, 4},
					{7, 2, 1},
				}, 0, 0, 3, 1,
			},
			output: output{0, errors.New("invalid start or end coordinates")},
		},
		{
			name: "異常時: girdが正方形でない場合、エラーを返すこと",
			input: input{
				[][]int{
					{3, 7, 1},
					{6, 3, 4},
					{7, 2},
				}, 0, 0, 2, 1,
			},
			output: output{0, errors.New("invalid grid length")},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result, err := CalculateOptimalPath(c.input.grid, c.input.startX, c.input.startY, c.input.endX, c.input.endY)
			if result != c.output.totalCost {
				t.Errorf("FindCalculateOptimalPath(%v, %v, %v, %v, %v) result %v, %v expected %v, %v", c.input.grid, c.input.startX, c.input.startY, c.input.endX, c.input.endY, result, err, c.output.totalCost, c.output.err)
				return
			}
			if c.output.err == nil {
				if err != nil {
					t.Errorf("FindCalculateOptimalPath(%v, %v, %v, %v, %v) result %v, %v expected %v, %v", c.input.grid, c.input.startX, c.input.startY, c.input.endX, c.input.endY, result, err, c.output.totalCost, c.output.err)
				}
			} else {
				if err != nil && err.Error() != c.output.err.Error() {
					t.Errorf("FindCalculateOptimalPath(%v, %v, %v, %v, %v) result %v, %v expected %v, %v", c.input.grid, c.input.startX, c.input.startY, c.input.endX, c.input.endY, result, err, c.output.totalCost, c.output.err)
				}
			}
		})
	}
}
