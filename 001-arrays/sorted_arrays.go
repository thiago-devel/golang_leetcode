package main

import "math"

/*
	ex: [-7, -3, 2, 3, 11]

	sublist A direction
	   >   >   >  >  >

	  [-7, -3, 2, 3, 11]

	sublist B direction
	   <   <   <  <  <

	loopSize = 5
	create a resultList with 5 positions of zeros'
ç	sublist A selector Index = starting with index 0
	sublist B selector Index = starting with index equals to (loopSize - 1) 4

	loop 1:
		position = (loopSize - 1) = (5 - 1) = 4
		position is equal or greater than zero so...
		a = absolute number of the item at sublistAselectorIndex (nums[0]) = (| -7 | = 7 ) = 7
		b = absolute number of the item at sublistBselectorIndex (nums[4]) = (| +11 | = 11 ) = 11
		test if 'a' (49) major than 'b' (121) result false
		so...
			on position (4) on resultList
			put the value of 'b' which is 121
			state: [0 , 0, 0, 0, 121]
			...then decrement 1 from sublistBselectorIndex (3)
		decrement 1 from position (3)
	loop 2:
		position = 3
		position is equal or greater than zero so...
		a = absolute number of the item at sublistAselectorIndex (nums[0]) = (| -7 | = 7 ) = 7
		b = absolute number of the item at sublistBselectorIndex (nums[3]) = (| +3 | = 3 ) = 3
		test if 'a' (49) major than 'b' (9) result true
		so...
			on position (3) of resultList
			put the valeus square of 'a' which is 49
			state: [0 , 0, 0, 49, 121]
			...then increment 1 from sublistAselectorIndex (1)
		decrement 1 from position (2)
	loop 3:
		position = 2
		position is equal or greater than zero so...
		a = absolute number of the item at sublistAselectorIndex (nums[1]) = (| -3 | = 3 ) = 3
		b = absolute number of the item at sublistBselectorIndex (nums[3]) = (| +3 | = 3 ) = 3
		test if 'a' (9) major than 'b' (9) result false
		so...
			on position (2) of resultList
			put the value of 'b' which is 9
			state: [0 , 0, 9, 49, 121]
			...then decrement 1 from sublistBselectorIndex (2)
		decrement 1 from position (1)
	loop 4:
		position = 1
		position is equal or greater than zero so...
		a = absolute number of the item at sublistAselectorIndex (nums[1]) = (|-3 | = 3 ) = 3
		b = absolute number of the item at sublistBselectorIndex (nums[2]) = (| 2 | = 2 ) = 2
		test if 'a' (9) major than 'b' (4) result true
		so...
			on position (1) of resultList
			put the value of 'a' which is 9
			state: [0 , 9, 9, 49, 121]
			...then increment 1 from sublistAselectorIndex (2)
		decrement 1 from position (0)
	loop 5:
		position = 0
		position is equal or greater than zero so...
		a = absolute number of the item at sublistAselectorIndex (nums[2]) = (| 2 | = 2 ) = 2
		b = absolute number of the item at sublistBselectorIndex (nums[2]) = (| 2 | = 2 ) = 2
		test if 'a' (4) major than 'b' (4) result false
		so...
			on position (0) of resultList
			put the value of 'b' which is 4
			state: [4 , 9, 9, 49, 121]
			...then decrement 1 from sublistBselectorIndex (1)
		decrement 1 from position (-1)

	end of loop
	The solution was founded interacting with the array just 1 time [ O(n) solution ]

*/
func SquareAndSortNumbers(nums []int) []int {
	var loopSize int = len(nums)
	// create the result array based on the initial
	// size of (slots) in the nums arrays
	resultList := make([]int, loopSize)

	sublistAselectorIndex := 0
	sublistBselectorIndex := loopSize - 1

	for position := (loopSize - 1); position >= 0; position-- {
		a := math.Abs(float64(nums[sublistAselectorIndex]))
		b := math.Abs(float64(nums[sublistBselectorIndex]))
		if a > b {
			resultList[position] = nums[sublistAselectorIndex] * nums[sublistAselectorIndex]
			sublistAselectorIndex++
		} else {
			resultList[position] = nums[sublistBselectorIndex] * nums[sublistBselectorIndex]
			sublistBselectorIndex--
		}
	}

	return resultList
}
