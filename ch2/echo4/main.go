// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 33.
//!+

// Echo4 prints its command-line arguments.
package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline") //go run ./ch2/echo4/main.go -n  传了n 值为true 没传默认是false
var sep = flag.String("s", " ", "separator")

//go run ./ch2/echo4/main.go a bv
//go run ./ch2/echo4/main.go -n a bv
//go run ./ch2/echo4/main.go -s / a bv
//go run ./ch2/echo4/main.go -help
func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}

//!-
