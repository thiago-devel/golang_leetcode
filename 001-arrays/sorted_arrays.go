package main

import "leetcode/001-arrays/utils"

// Given an integer array nums sorted in non-decreasing
// order, return an array of the squares of each number
// sorted in non-decreasing order.
func SquareAndSortNumbers(nums []int) []int {
	numsLimit := nums
	if len(nums) > 10000 {
		numsLimit = nums[0:10000]
	}

	return sortNonDecreassing(squareNumbers(numsLimit))
}

func squareNumbers(nums []int) []int {
	squares := []int{}
	for _, number := range nums {
		square := (number * number)
		squares = append(squares, square)
	}
	return squares
}

func sortNonDecreassing(nums []int) []int {
	return utils.NumbersInsertionSort(nums)
}
