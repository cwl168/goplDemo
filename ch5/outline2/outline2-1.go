// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 133.

// Outline prints the outline of an HTML document tree.
package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

//go run ch5/outline2/outline2-1.go https://www.baidu.com
//练习5.12： gopl.io/ch5/outline2（5.5节）的startElement和endElement共用了全局变量depth，将它们修改为匿名函数，使其共享outline中的局部变量
func main() {
	for _, url := range os.Args[1:] {
		outline(url)
	}
}

func outline(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}
	//使用匿名函数实现
	var depth int
	var startElement func(n *html.Node)
	var endElement func(n *html.Node)
	startElement = func(n *html.Node) {
		if n.Type == html.ElementNode {
			attr := ""
			for _, a := range n.Attr {
				attr += " " + a.Key + "=" + "\"" + a.Val + "\" "
			}
			fmt.Printf("%*s<%s%s", depth*2, "", n.Data, attr)
			depth++
		}
		if n.Type == html.ElementNode && n.FirstChild == nil && n.Data != "script" {
			fmt.Printf("/>\n")
		} else if n.Type == html.ElementNode {
			fmt.Printf(">\n")
		}

		if n.Type == html.TextNode {
			fmt.Printf("%*s %s\n", depth*2, "", n.Data)
		}
	}
	endElement = func(n *html.Node) {
		if n.Type == html.ElementNode && n.FirstChild == nil && n.Data != "script" {
			depth--
			fmt.Printf("\n")
			return
		}
		if n.Type == html.ElementNode {
			depth--

			fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
		}
	}
	//!+call
	forEachNode(doc, startElement, endElement)
	//!-call

	return nil
}

//!+forEachNode
// forEachNode calls the functions pre(x) and post(x) for each node
// x in the tree rooted at n. Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder).
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

/**
<html>
<head>
	<script>
		location.replace(location.href.replace("https://","http://"));
	</script>
</head>
<body>
	<noscript><meta http-equiv="refresh" content="0;url=http://www.baidu.com/"></noscript>
</body>
</html>

*/
