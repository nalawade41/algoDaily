package main

import "fmt"

// Problem Link - https://www.hackerrank.com/challenges/string-construction/problem

func main() {
	s1 := "wouldyoulikefries"
	s2 := "abcabcabcabcabcabc"
	fmt.Println(twoStrings(s1,s2))
}

func twoStrings(s1 string, s2 string) string {
	values := make([]bool, 26)

	for _, v := range s1 {
		values[v-'a'] = true
	}

	for _, v := range s2 {
		if values[v-'a'] {
			return "YES"
		}
	}
	return "NO"

}
