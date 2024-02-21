package main

import (
	"testing"
)

func TestAdd(t *testing.T) {

	testCases := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"Add positive numbers", 2, 3, 5},
		{"Add negative numbers", -1, -2, -3},
		{"Add zero", 0, 0, 0},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := Add(tc.a, tc.b)
			expectedResult := tc.expected
			if result != expectedResult {
				t.Errorf("Add(%d, %d) = %d is wrong, correct is %d",
					tc.a, tc.b, result, expectedResult)
			}
		})
	}
}

func TestFactorial(t *testing.T) {
	testCases := []struct {
		name     string
		n        int
		expected int
	}{
		{"Factorial of 0", 0, 1},
		{"Factorial of 1", 1, 1},
		{"Factorial of 5", 5, 120},
		{"Factorial of negative number", -1, 1},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := Factorial(tc.n)
			if result != tc.expected {
				t.Errorf("Test %s failed. Expected %d, got %d", tc.name, tc.expected, result)
			}
		})
	}
}
