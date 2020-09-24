// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 153.

// Title3 prints the title of an HTML document specified by a URL.
package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

// Copied from gopl.io/ch5/outline2.
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

//!+
// soleTitle returns the text of the first non-empty title element
// in doc, and an error if there was not exactly one.
//如果检测到有多个<title>，会调用panic，阻止函数继续递归， 并将特殊类型bailout作为panic的参数。
func soleTitle(doc *html.Node) (title string, err error) {
	type bailout struct{}

	defer func() {
		switch p := recover(); p {
		case nil:
			// no panic
		case bailout{}:
			// "expected" panic
			err = fmt.Errorf("multiple title elements")
		default:
			panic(p) // unexpected panic; carry on panicking
		}
	}()

	// Bail out of recursion if we find more than one non-empty title.
	// 如果我们发现不止一个非空标题，就退出递归。
	forEachNode(doc, func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
			//大于1个title 直接panic
			if title != "" {
				panic(bailout{}) // multiple title elements
			}
			title = n.FirstChild.Data
		}
	}, nil)
	if title == "" {
		return "", fmt.Errorf("no title element")
	}
	return title, nil //或者直接return
}

//!-

func title() error {
	/*resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	// Check Content-Type is HTML (e.g., "text/html; charset=utf-8").
	ct := resp.Header.Get("Content-Type")
	if ct != "text/html" && !strings.HasPrefix(ct, "text/html;") {
		return fmt.Errorf("%s has type %s, not text/html", url, ct)
	}
	doc, err := html.Parse(resp.Body)
	*/

	doc, err := html.Parse(os.Stdin) //必须是html内容
	if err != nil {
		return fmt.Errorf("parsing err: %v", err)
	}
	title, err := soleTitle(doc)
	if err != nil {
		return err
	}
	fmt.Println(title)
	return nil
}

//如果HTML页面包含多个<title>，该函数会给调用者返回一个错误 （error）。在soleTitle内部处理时，如果检测到有多个<title>，会调用panic，阻止函数继续递归， 并将特殊类型bailout作为panic的参数。
//cat ch5/title3/index.html | go run ch5/title3/title1.go
func main() {
	//for _, arg := range os.Args[1:] {
	//	if err := title(arg); err != nil {
	//		fmt.Fprintf(os.Stderr, "title: %v\n", err)
	//	}
	//}
	if err := title(); err != nil {
		fmt.Fprintf(os.Stderr, "title: %v\n", err)
	}
}
