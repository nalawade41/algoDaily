package main

import "testing"


type TestDataItem struct {
	input int
	expected int
}

func TestGetBinaryGap(t *testing.T) {
	values := []TestDataItem{
		{input: 9,expected:2 },
		{input: 529,expected:4},
		{input: 20,expected:1},
		{input: 15,expected:0},
		{input: 32,expected:0},
		{input:1041,expected:5},
	}

	for _, testItem := range values {
		result := getBinaryGap(testItem.input)

		if result != testItem.expected {
			t.Errorf("getBinaryGap() failed expected %v, got %v",testItem.expected, result)
		}
	}

}
