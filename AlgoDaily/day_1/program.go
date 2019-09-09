package main

import "fmt"

func main() {
	nums1 := []int {4,17,4,4,15,16,17,6,7}
	nums2 := []int {15,2,6,20,17,17,8,4,5}
	fmt.Println(intersection(nums1, nums2))
}

func intersection (nums1 []int, nums2 []int) []int {
	values := make(map[int]int)

	for _, v := range nums1 {
		values[v] = 1
	}

	for _, v := range nums2 {
		if c, ok := values[v]; ok {
			values[v] = c+1
		}
	}

	var output []int
	for i, v := range values {
		if v > 1 {
			output = append(output, i)
		}
	}

	return output
}