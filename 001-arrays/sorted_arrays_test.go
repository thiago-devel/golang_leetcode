package main

import (
	"leetcode/001-arrays/utils"
	"testing"
)

func TestGenerateSquare(t *testing.T) {
	arrayTest := []int{5, 4, 3, 2, 1}
	want := []int{1, 4, 9, 16, 25}

	got := SquareAndSortNumbers(arrayTest)

	if !utils.Equal(got, want) {
		t.Errorf("got %v wanted %v", got, want)
	}

}

func TestGenerateSquare2(t *testing.T) {
	arrayTest := utils.MakeRange(1, 5)
	want := []int{1, 4, 9, 16, 25}

	got := SquareAndSortNumbers(arrayTest)

	if !utils.Equal(got, want) {
		t.Errorf("got %v wanted %v", got, want)
	}

}

func TestSquareAndSortNumbers(t *testing.T) {
	arrayTest := []int{5, 4, 3, 2, 1}
	want := []int{1, 4, 9, 16, 25}

	got := SquareAndSortNumbers(arrayTest)

	if !utils.Equal(got, want) {
		t.Errorf("got %v wanted %v", got, want)
	}

}

func TestSquareAndSortNumbers2(t *testing.T) {
	arrayTest := []int{-4, -1, 0, 3, 10}
	want := []int{0, 1, 9, 16, 100}

	got := SquareAndSortNumbers(arrayTest)

	if !utils.Equal(got, want) {
		t.Errorf("got %v wanted %v", got, want)
	}

}

func TestSquareAndSortNumbers3(t *testing.T) {
	arrayTest := []int{-7, -3, 2, 3, 11}
	want := []int{4, 9, 9, 49, 121}

	got := SquareAndSortNumbers(arrayTest)

	if !utils.Equal(got, want) {
		t.Errorf("got %v wanted %v", got, want)
	}

}
