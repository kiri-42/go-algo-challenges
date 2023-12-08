package main

import (
	"errors"
	"fmt"
)

func FindMaxValue(numbers []int) (maxValue int, err error) {
	if len(numbers) == 0 {
		return 0, errors.New("empty slice")
	}

	maxValue = -9999
	for _, num := range numbers {
		if num < -1000 || num > 1000 {
			return 0, errors.New("invalid value in slice")
		}

		if num > maxValue {
			maxValue = num
		}
	}
	return maxValue, nil
}

func main() {
	maxValue, _ := FindMaxValue([]int{2, 3})
	fmt.Println(maxValue)
}
