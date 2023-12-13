package main

import "errors"

func FindMinAndIndex(nums []int) (min int, index int, err error) {
	if err = Validate(nums); err != nil {
		return 0, 0, err
	}

	min = nums[0]
	for i, num := range nums {
		if num < min {
			min = num
			index = i
		}
	}

	return min, index, nil
}

func Validate(nums []int) error {
	if len(nums) == 0 {
		return errors.New("invalid input: slice cannot be empty")
	}

	return nil
}
