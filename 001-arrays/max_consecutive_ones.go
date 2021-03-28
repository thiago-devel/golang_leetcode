package main

// find maximum consecutive ones on a given positive number array
// nums only accepts values Zeros or Ones as values
// will return the number of maximum consecutives ones in the nums arrays
func FindMaxConsecutiveOnes(nums []uint16) int {
	// if the array if empty, return 0
	if len(nums) < 1 {
		return 0
	}

	var maxConsecutivesOnes int = 0

	oneWindowSize := 0

	// starts to search for 'one's windows' inside of the list
	for i := 0; i < len(nums); i++ {

		if nums[i] == 1 {
			// starting a new 'one window'...

			// calculating how many Ones have int that Window
			oneWindowSize++

		} else if ((i + 1) == len(nums)) || (nums[i] == 0) {
			// found the end of current 'one window'....
			// reset the one window size counter for the next iteration
			oneWindowSize = 0
		}

		if oneWindowSize > maxConsecutivesOnes {
			// if the one window size is grater then maximum concecutive ones...
			// increment max
			maxConsecutivesOnes = oneWindowSize
		}
	}

	return maxConsecutivesOnes
}
