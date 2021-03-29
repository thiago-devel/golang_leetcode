package utils

// Reference and credits: https://stackoverflow.com/a/39868255/5003355
func MakeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

// Equal tells whether a and b contain the same elements.
// A nil argument is equivalent to an empty slice.
// Reference and credits : https://yourbasic.org/golang/compare-slices/
func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

// reference: https://stackabuse.com/insertion-sort-in-javascript/
// TODO: explain
func NumbersInsertionSort(nums []int) []int {
	sortedList := append([]int(nil), nums...)

	for i := 1; i < len(sortedList); i++ {

		//
		currentNumber := sortedList[i]

		//
		j := i - 1

		//
		for (j > -1) && (currentNumber < sortedList[j]) {
			//
			sortedList[(j + 1)] = sortedList[j]

			//
			j--
		}

		//
		sortedList[(j + 1)] = currentNumber
	}

	return sortedList
}
