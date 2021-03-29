package main

import (
	"leetcode/001-arrays/utils"
	"testing"
)

func TestFind1EvenNumberDigitsInListFaced(t *testing.T) {
	binaryList := []int{555, 901, 482, 1771}
	var want int = 1

	got := FindNumbersForced(binaryList)

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestFind2EvenNumberDigitsInListForced(t *testing.T) {
	binaryList := []int{12, 345, 2, 6, 7896}
	var want int = 2

	got := FindNumbersForced(binaryList)

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestErrorToFindEvenNumberDigitsInListForced(t *testing.T) {
	binaryList := []int{2, 444, 66666, 8888888}
	var want int = 1

	got := FindNumbersForced(binaryList)

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestErrorToFindEvenNumbersDigitsInListForced2(t *testing.T) {
	binaryList := []int{}
	var want int = 0

	got := FindNumbersForced(binaryList)

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestErrorToFindEvenNumbersDigitsInListForced3(t *testing.T) {
	var want int = 1

	got := FindNumbersForced(utils.MakeRange(501, 1002))

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
