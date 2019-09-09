package main

import "testing"

type TestCases struct {
	first string
	expected int
}
func TestStringConstruction (t *testing.T){
	cases := []TestCases{
		{first:"scfg", expected:4},
		{first:"bccb", expected:2},
	}

	for _, tcase := range cases {
		result := stringConstruction(tcase.first)

		if result != tcase.expected {
			t.Errorf("stringConstruction() failed expected %v, got %v",tcase.expected, result)
		}
	}
}
