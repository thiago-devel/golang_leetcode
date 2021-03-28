package main

// Given an array nums of integers, return
// how many of them contain an even number of digits
func FindNumbersForced(nums []int) int {
	result := 0

	numsSize := len(nums)
	if numsSize < 1 {
		return 0
	}

	if numsSize > 500 {
		nums = nums[0:500]
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
			number = 100000
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

	return result
}
