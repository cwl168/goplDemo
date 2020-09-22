// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 130.

// The wait program waits for an HTTP server to start responding.
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

//!+
// WaitForServer attempts to contact the server of a URL.
// It tries for one minute using exponential back-off.
// It reports an error if all attempts fail.
func WaitForServer(url string) error {
	const timeout = 20 * time.Minute
	deadline := time.Now().Add(timeout)
	//退避重试算法 (重试等待时间指数级增长)
	//偶然错误，进行重试。 如果错误的发生是偶然性的，或由不可预知的问题导致的。一个明智的选择是重新尝试失败的操作。在重试时，我们需要限制重试的时间间隔或重试的次数，防止无限制的重试。
	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url)
		if err == nil {
			return nil // success
		}
		log.Printf("server not responding (%s); retrying...", err)
		fmt.Println(time.Second, time.Second<<uint(tries), uint(tries)) //左移相当于乘以2  //等待时间随着以二为底的指数增长。如果重试失败，那么下次的等待时间将会是上次的等待时间二倍。如果重试次数大于最大重试次数，那么包将从包队列中去除
		time.Sleep(time.Second << uint(tries))                          // exponential back-off
	}
	return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}

//!-
//go run ch5/wait/wait.go https://www.google.com.hk/
func main() {
	/*if err := Ping(); err != nil {
		log.Printf("ping failed: %v; networking disabled", err)
	}*/

	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: wait url\n")
		os.Exit(1)
	}
	url := os.Args[1]
	//!+main
	// (In function main.)
	//程序无法继续运行，输出错误信息并结束程序。
	if err := WaitForServer(url); err != nil {
		fmt.Fprintf(os.Stderr, "Site is down: %v\n", err)
		os.Exit(1)
	}
	//!-main
}
