package main

import (
	"fmt"
	"strconv"
	"strings"
)

func SumNumbers(numbers string) (sum int, err error) {
	if numbers == "" {
		return 0, nil
	}

	numberList := strings.Split(numbers, ",")

	for _, number := range numberList {
		numberInt, err := strconv.Atoi(number)
		if err != nil {
			return 0, fmt.Errorf("invalid number: %v", number)
		}

		sum += numberInt
	}

	return sum, nil
}
