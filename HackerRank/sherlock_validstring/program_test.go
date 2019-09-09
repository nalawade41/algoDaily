package main

import "testing"

type TestCase struct {
	input string
	expected string
}

func TestIsValid(t *testing.T) {
	cases := []TestCase {
		{input:"aabbcd", expected:"NO"},
		{input:"aabbccddeefghi", expected:"NO"},
		{input:"abcdefghhgfedecba", expected:"YES"},
	}

	for _, tcase := range cases {
		result := isValid(tcase.input)

		if result != tcase.expected {
			t.Errorf("TestIsValid() failed expected %v, got %v",tcase.expected, result)
		}
	}
}