// Copyright © 2017 xingdl2007@gmail.com
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// 1.3 measure the time of different implementation of echo
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func echo1() string {
	start := time.Now()
	defer func() {
		fmt.Printf("echo1: %v ns\n", time.Since(start).Nanoseconds())
	}()

	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	return s
}

func echo2() string {
	start := time.Now()
	defer func() {
		fmt.Printf("echo2: %v ns\n", time.Since(start).Nanoseconds())
	}()

	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	return s
}

func echo3() string {
	start := time.Now()
	defer func() {
		fmt.Printf("echo3: %v ns\n", time.Since(start).Nanoseconds())
	}()

	return strings.Join(os.Args[1:], " ")
}

//做实验测量潜在低效的版本和使用了strings.Join的版本的运行时间差异。（1.6节讲解了部分time包，11.4节展示了如何写标准测试程序，以得到系统性的性能评测。）
//go run ch1/practice/1.3.go a b c d e f g h i j k l m n o p q r s t u v w x y z
/**
echo1: 5375 ns
echo2: 1888 ns
echo3: 672 ns

*/
func main() {
	echo1()
	echo2()
	echo3()
}
