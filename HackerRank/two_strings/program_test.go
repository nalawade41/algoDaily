package main

import "testing"

type TestCases struct {
	first string
	second string
	expected string
}
func TestTwoString(t *testing.T) {
	cases := []TestCases{
		{first:"wouldyoulikefries", second:"abcabcabcabcabcabc", expected:"NO"},
		{first:"hackerrankcommunity", second:"cdecdecdecde", expected:"YES"},
		{first:"jackandjill", second:"wentupthehill", expected:"YES"},
		{first:"writetoyourparents", second:"fghmqzldbc", expected:"NO"},
	}

	for _, tcase := range cases {
		result := twoStrings(tcase.first, tcase.second)

		if result != tcase.expected {
			t.Errorf("TestTwoString() failed expected %v, got %v",tcase.expected, result)
		}
	}
}
