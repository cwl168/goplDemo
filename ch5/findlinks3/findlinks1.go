// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 139.

// Findlinks3 crawls the web, starting with the URLs on the command line.
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"

	"gopl.io/ch5/links"
)

var sum int

//!+breadthFirst
// breadthFirst calls f for each item in the worklist.
// Any items returned by f are added to the worklist.
// f is called at most once for each item.
//广度优先算法。
func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

/*
抓取页面的所有连接
*/
func crawl(url string) []string {
	sum++
	go save(url)
	fmt.Printf("%d|%s\n", sum, url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

/*
保存页面到文件
*/
func save(u string) bool {
	urlObj, _ := url.Parse(u)
	path := "/Users/caoweilin/crawl/" + urlObj.Host
	if urlObj.Path == "" || urlObj.Path == "/" {
		urlObj.Path = "/index.html"
	}
	filename := path + urlObj.Path //重点注意文件名
	fmt.Println(path + "-------" + filename + "------------" + u)
	//打开文件
	f, _ := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0755)
	//读取链接
	resp, geterr := http.Get(u)

	if geterr != nil || resp.StatusCode != http.StatusOK {
		//resp.Body.Close()
		return false
	}
	body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println(body)
	//创建保存目录
	_, err := os.Stat(path)
	if err != nil {
		os.MkdirAll(path, 0755)
	}
	io.WriteString(f, string(body))
	resp.Body.Close()
	body = nil
	return true
}

/*func save2(u string) bool {
	urlObj, _ := url.Parse(u)
	path := "/Users/caoweilin/crawl/" + urlObj.Host
	if urlObj.Path == "" || urlObj.Path == "/" {
		urlObj.Path = "/index.html"
	}
	filename := path + urlObj.Path //重点注意文件名
	fmt.Println("filename:" + filename)
	fmt.Println("path:" + path)
	//打开文件
	f, ferr := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0755)
	if ferr != nil {
		fmt.Fprintf(os.Stderr, "os.OpenFile err: %v\n", ferr)
		os.Exit(1)
	}
	//读取链接
	resp, geterr := http.Get(u)

	if geterr != nil || resp.StatusCode != http.StatusOK {
		//resp.Body.Close()
		return false
	}
	body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println(body)
	//创建保存目录
	_, serr := os.Stat(path)
	if serr != nil {
		merr := os.MkdirAll(path, 0755)
		if merr != nil {
			fmt.Fprintf(os.Stderr, "os.MkdirAll err: %v\n", merr)
			os.Exit(1)
		}
	}

	io.WriteString(f, string(body))
	resp.Body.Close()
	body = nil
	return true
}*/

//!+main
//go run ch5/findlinks3/findlinks1.go http://www.baidu.com     https的百度域名无法抓取
//修改crawl，使其能保存发现的页面，必要时，可以创建目录来保存这些页面。只保存来自 原始域名下的页面。假设初始页面在golang.org下，就不要保存vimeo.com下的页面。
func main() {
	// Crawl the web breadth-first,
	// starting from the command-line arguments.
	breadthFirst(crawl, os.Args[1:])
}

//!-main
