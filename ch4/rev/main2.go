package main

import (
	"fmt"
)

func main() {
	testArr := [5]int{0, 1, 2, 3, 4}
	reverse2(&testArr)
	fmt.Println(testArr)
}

func reverse1(s *[5]int) {
	i, j := 0, len(*s)-1
	for i < j {
		(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
		i += 1
		j -= 1
	}
}
func reverse2(s *[5]int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
