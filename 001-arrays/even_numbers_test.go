package main

import (
	"leetcode/001-arrays/utils"
	"testing"
)

func TestFind1EvenNumberDigitsInList(t *testing.T) {
	binaryList := []int{555, 901, 482, 1771}
	var want int = 1

	got, _ := FindNumbers(binaryList)

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestFind2EvenNumberDigitsInList(t *testing.T) {
	binaryList := []int{12, 345, 2, 6, 7896}
	var want int = 2

	got, _ := FindNumbers(binaryList)

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestErrorToFindEvenNumberDigitsInList(t *testing.T) {
	binaryList := []int{2, 444, 66666, 8888888}

	got, errorObj := FindNumbers(binaryList)

	if errorObj == nil {
		t.Errorf("Invalid array element Error was expected. Got %d", got)
	}
}

func TestErrorToFindEvenNumbersDigitsInList2(t *testing.T) {
	binaryList := []int{}

	got, errorObj := FindNumbers(binaryList)

	if errorObj == nil {
		t.Errorf("Invalid array range Error was expected. Got %d", got)
	}
}

func TestErrorToFindEvenNumbersDigitsInList3(t *testing.T) {

	got, errorObj := FindNumbers(utils.MakeRange(1, 501))

	if errorObj == nil {
		t.Errorf("Invalid array range Error was expected. Got %d", got)
	}
}
