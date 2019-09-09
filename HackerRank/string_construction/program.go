package main

import "fmt"

// Problem Link = https://www.hackerrank.com/challenges/two-strings/problem

func main() {
	fmt.Println(stringConstruction("abcda"))
}

func stringConstruction(s string) int {
	cost := make([]int,26)
	for _, v := range s {
		cost[v-'a'] = 1
	}

	totalCost := 0
	for _, v := range cost {
		totalCost+= v
	}

	return totalCost
}
