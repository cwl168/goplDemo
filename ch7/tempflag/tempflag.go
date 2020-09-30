// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 181.

// Tempflag prints the value of its -temp (temperature) flag.
package main

import (
	"flag"
	"fmt"

	"gopl.io/ch7/tempconv"
)

//!+
var temp = tempconv.CelsiusFlag("temp", 20.0, "the temperature")

//go run ch7/tempflag/tempflag.go
// go run ch7/tempflag/tempflag.go -temp -18C
//go run ch7/tempflag/tempflag.go -help
//go run ch7/tempflag/tempflag.go -temp 212°F
func main() {
	flag.Parse()
	fmt.Println(*temp)
}

//!-
