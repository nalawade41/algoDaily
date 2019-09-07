package main

import "fmt"

func main() {

	fmt.Println(reverse("thisistest"))
	fmt.Println(optimizedReverse("thisistest"))
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