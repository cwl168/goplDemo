// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 74.

// Printints demonstrates the use of bytes.Buffer to format a string.
package main

import (
	"bytes"
	"fmt"
)

//!+
// intsToString is like fmt.Sprint(values) but adds commas.
//将切片转化为字符串
func intsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[') //写入byte
	for i, v := range values {
		fmt.Println(i, v)
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%dr", v)
	}
	buf.WriteByte(']')
	return buf.String()
}

func main() {
	fmt.Printf("%T\n", []int{1, 2, 3}) // "[1, 2, 3]"
	s := intsToString([]int{1, 2, 3})
	fmt.Println(s) // "[1, 2, 3]"
	fmt.Printf("%T\n", s)
}

//!-
