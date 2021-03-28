package main

import "testing"

func TestMaxConsecutiveOnesValid(t *testing.T) {
	binaryList := []uint16{1, 1, 0, 1, 1, 1}
	var want int = 3

	got := FindMaxConsecutiveOnes(binaryList)

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
func TestMaxConsecutiveOnesValid2(t *testing.T) {
	binaryList := []uint16{0, 1, 1, 1, 1, 0, 1, 1, 1, 0}
	var want int = 4

	got := FindMaxConsecutiveOnes(binaryList)

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
func TestMaxConsecutiveOnesEmpty(t *testing.T) {
	binaryList := []uint16{}
	var want int = 0

	got := FindMaxConsecutiveOnes(binaryList)

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
func TestMaxConsecutiveOnesZeroMatches(t *testing.T) {
	binaryList := []uint16{0, 0, 0, 0, 0, 0}
	var want int = 0

	got := FindMaxConsecutiveOnes(binaryList)

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
