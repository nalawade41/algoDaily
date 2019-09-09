package main

import "testing"

type TestCase struct {
	first []int
	second []int
	expected []int
}

func TestIntersection(t *testing.T) {
	testCases := []TestCase {
		{first:[]int{6,0,12,10,16}, second:[]int{3,15,18,20,15}, expected:[]int{}},
		{first:[]int{1,5,2,12,6}, second:[]int{13,10,9,5,8}, expected:[]int{5}},
		{first:[]int{4,17,4,4,15,16,17,6,7}, second:[]int{15,2,6,20,17,17,8,4,5	}, expected:[]int{15,6,17,4}},
		{first:[]int{3}, second:[]int{15}, expected:[]int{}},
		{first:[]int{2,16,8,9}, second:[]int{14,15,2,20}, expected:[]int{2}},
	}

	for _, test := range testCases {
		result := intersection(test.first, test.second)

		if len(result) != len(test.expected) {
			t.Errorf("TestIntersection() failed expected %v, got %v",test.expected, result)
		}
	}
}
