package math

import "testing"

type AddResult struct {
	x              int
	y              int
	expectedOutput int
}

var addResults = []AddResult{
	{1, 1, 2},
	{3, 8, 11},
	{-2, -2, -4},
}

func Test1Plus1(test *testing.T) {
	if Add(1, 1) != 2 {
		test.Error("1 + 1 != 2")
	}
}

func Test3Plus8(test *testing.T) {
	if Add(3, 8) != 11 {
		test.Error("3 + 8 != 11")
	}
}

func TestAdd(test *testing.T) {
	for _, addTest := range addResults {
		result := Add(addTest.x, addTest.y)

		if result != addTest.expectedOutput {
			test.Errorf("Add(%d, %d) Expected: %d; Got: %d", addTest.x, addTest.y, addTest.expectedOutput, result)
		}
	}
}
