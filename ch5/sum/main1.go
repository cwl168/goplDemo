// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 142.

// The sum program demonstrates a variadic function.
package main

import (
	"fmt"
)

func max(vals ...int) (interface{}, error) {
	num := len(vals)
	if num == 0 {
		return 0, fmt.Errorf("max: %s", "至少传递一个参数")
	}
	//max := vals[:1][0]
	max := vals[0]
	for _, val := range vals {
		if val > max {
			max = val
		}
	}
	return max, nil
}
func main() {
	fmt.Println(max())
	fmt.Println(max(1, 2, 3, 4)) //  "10"
}
