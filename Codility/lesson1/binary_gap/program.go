package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter vale: ")
	s, _ := reader.ReadString('\n')
	n, err := strconv.Atoi(s)
	if err != nil {
		fmt.Errorf("You need enter valid integer value")
	}
	fmt.Println(getBinaryGap(n))
}

func getBinaryGap(n int) int {
	var shouldSart bool
	var count int
	var highestCount int
	for n > 0 {
		b := n % 2
		n = n/2
		if b == 1 {
			shouldSart = true
		}

		if shouldSart && b == 0 {
			count ++
		} else {
			if highestCount < count {
				highestCount = count
			}
			count = 0
		}
	}
	if highestCount < count {
		highestCount = count
	}
	return highestCount
}