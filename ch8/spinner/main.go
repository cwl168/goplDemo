// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 218.

// Spinner displays an animation while computing the 45th Fibonacci number.
package main

import (
	"fmt"
	"runtime"
	"time"
)

//!+
//计算菲波那契数列的第45个元素值
func main() {
	defer func() {
		time.Sleep(time.Second * 2)
		fmt.Println("\nthe number of goroutines: ", runtime.NumGoroutine())
	}()
	go spinner(100 * time.Millisecond)
	const n = 45
	fibN := fib(n) // slow
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}

//当主函数返回时，所有的goroutine都会直接打断，程序退出。但是如果是常驻内存，会有goroutine泄露,该goroutine不会退出
func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}

//!-
