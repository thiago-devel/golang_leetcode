package main

import (
	"errors"
	"fmt"
)

// Given an array nums of integers, return
// how many of them contain an even number of digits
func FindNumbers(nums []int) (int, error) {
	result := 0

	numsSize := len(nums)
	if numsSize < 1 || numsSize > 500 {
		return 0, errors.New(fmt.Sprintf("Invalid array range: %d", numsSize))
	}

	/*
		for _, number := range nums {

			test := 0
			rest := number

			for {
				test++
				rest = rest / 10

				if rest == 0 {
					break
				}
			}

			if test%2 == 0 {
				result++
			}
		}
	*/

	// old style
	for i := 0; i < len(nums); i++ {

		number := nums[i]

		if number > 100000 {
			return 0, errors.New(fmt.Sprintf("Invalid array element: %d", numsSize))
		}

		test := 0
		rest := number

		for {
			test++
			rest = rest / 10

			if rest == 0 {
				break
			}
		}

		if (test % 2) == 0 {
			result++
		}
	}

	return result, nil
}
