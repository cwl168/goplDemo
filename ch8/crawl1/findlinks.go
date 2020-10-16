// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 240.

// Crawl1 crawls web links starting with the command-line arguments.
//
// This version quickly exhausts available file descriptors
// due to excessive concurrent calls to links.Extract.
//
// Also, it never terminates because the worklist is never closed.
package main

import (
	"fmt"
	"gopl.io/ch5/links"
	"log"
	"os"
)

//!+crawl
func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

//!-crawl

//!+main
//不过这一次我们用channel代替slice来做这个队 列。每一个对crawl的调用都会在他们自己的goroutine中进行并且会把他们抓到的链接发送回 worklist。
// 另外注意这里将命令行参数传入worklist也是在一个另外的goroutine中进行的，
func main() {
	worklist := make(chan []string)

	// Start with the command-line arguments.
	//go func() { worklist <- os.Args[1:] }()  //如果不在goroutine中接受参数，会阻塞
	worklist <- os.Args[1:]

	// Crawl the web concurrently.
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}
}

//!-main

/*
//!+output
$ go build gopl.io/ch8/crawl1
$ ./crawl1 http://gopl.io/
http://gopl.io/
https://golang.org/help/

https://golang.org/doc/
https://golang.org/blog/
...
2015/07/15 18:22:12 Get ...: dial tcp: lookup blog.golang.org: no such host
2015/07/15 18:22:12 Get ...: dial tcp 23.21.222.120:443: socket:
                                                        too many open files
...
//!-output
*/
