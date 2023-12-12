package main

import "errors"

func SumRange(numbers []int, start int, end int) (sum int, err error) {
	if err := Validate(numbers, start, end); err != nil {
		return 0, err
	}

	for i, number := range numbers {
		if i > end {
			return sum, nil
		}

		if i >= start {
			sum += number
		}
	}

	return sum, nil
}

func Validate(numbers []int, start int, end int) error {
	len := len(numbers)
	if len < 1 || len > 100 {
		return errors.New("error")
	}

	if start < 0 || start > len-1 {
		return errors.New("error")
	}

	if end < start || end > len-1 {
		return errors.New("error")
	}

	return nil
}
