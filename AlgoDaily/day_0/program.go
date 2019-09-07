package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter string to reverse: ")
	s, _ := reader.ReadString('\n')

	fmt.Println(reverse(s))
	fmt.Println(optimizedReverse(s))
}


func reverse(s string) string {
	rs:= ""
	i := len(s)-1
	for i >= 0 {
		rs+=string(s[i])
		i--
	}
	return rs
}

func optimizedReverse(s string) string {
	rs := []rune(s)
	i := len(s)-1
	j := 0
	for i>j {
		rs[i], rs[j] = rs[j], rs[i]
		i--
		j++
	}
	return string(rs)
}