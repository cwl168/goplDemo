package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	s5 := []int{1, 2, 3}
	fmt.Println(s5[1:])
	fmt.Println(s5[0:])
	fmt.Println(s5)
	fmt.Println(strings.Join(os.Args, " "))
}
