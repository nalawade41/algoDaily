package main

import (
	"fmt"
	"sort"
)

//Problem link = https://www.hackerrank.com/challenges/sherlock-and-valid-string/problem

func main() {
	fmt.Println(isValid("abcdefghhgfedecba"))
}

func isValid(s string) string {

	if len(s) <= 3 {
		return "YES"
	}

	freq := make([]int, 26)

	for _, v := range s {
		count := freq[v-'a']
		freq[v-'a'] = count+1
	}

	sort.Ints(freq)

	nonExistent := 0
	for freq[nonExistent] == 0 {
		nonExistent++
	}

	minFreq := freq[nonExistent]
	maxFreq := freq[25]

	if minFreq == maxFreq {
		return "YES"
	} else {
		if ((maxFreq-minFreq==1) && (maxFreq > freq[24])) ||
			((minFreq ==1) && freq[nonExistent+1] == maxFreq) {
			return "YES"
		}
	}
	return "NO"
}
