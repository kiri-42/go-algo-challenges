package main

import (
	"errors"
)

func FindEscapeCode(numbers []int) (code int, err error) {
	if err = Validate(numbers); err != nil {
		return 0, err
	}

	codeCount := map[int]int{}
	for _, number := range numbers {
		codeCount[number]++
	}

	max := 0
	for i, count := range codeCount {
		if count >= max && i > code {
			code = i
			max = count
		}
	}

	return code, nil
}

func Validate(numbers []int) error {
	if len(numbers) == 0 {
		return errors.New("no numbers provided")
	}

	for _, number := range numbers {
		if number < 0 || number > 9 {
			return errors.New("invalid number value")
		}
	}

	return nil
}
