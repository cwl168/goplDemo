// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 17.
//!+

// Fetchall fetches URLs in parallel and reports their times and sizes.
package main

import (
	"fmt"
	"net/http"
	"os"
	"runtime"
	"strings"
	"time"
)

//go run ch1/fetchall/main.go https://baidu.com https://sina.com.cn https://qq.com
//go run ch1/fetchall/main.go baidu.com sina.cn 163.com google.cn qq.com weibo.cn
func main() {
	defer func() {
		time.Sleep(time.Second)
		fmt.Printf("goroutine num %d\n", runtime.NumGoroutine())
	}()
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http") {
			url = "https://" + url
		}
		go fetch(url, ch) // start a goroutine
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel ch
	}
	/*for s := range ch {
		fmt.Println(s)
	}*/
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}
func fetch(url string, ch chan<- string) {
	/*start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)*/
	_,_ =http.Get(url)
	ch <- url
	//close(ch)
}

//!-
