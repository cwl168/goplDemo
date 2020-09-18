// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 122.
//!+main

// Findlinks1 prints the links in an HTML document read from standard input.
package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

//ch5/findlinks1/fetch https://www.sina.com.cn | go run ch5/findlinks1/main.go
//echo '<a href=www.baidu.com></a>' | go run ch5/findlinks1/main.go   获取网页内容中所有的href标签的链接
//cat ch5/findlinks1/index.html | go run ch5/findlinks1/main.go
//go run ch5/findlinks1/main.go  输入内容，然后ctrl+d
func main() {
	//input := bufio.NewScanner(os.Stdin)
	//input.Scan()
	//fmt.Println("你输入的是：", input.Text())
	//os.Exit(0)
	doc, err := html.Parse(os.Stdin) //必须是html内容
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	//链接地址不会换行
	fmt.Println(visit(nil, doc))
	//链接地址会换行
	//for _, link := range visit(nil, doc) {
	//	fmt.Println(link)
	//}
	count_map := map[string]int{}
	fmt.Println(count(count_map, doc))
	fmt.Println(visit3(nil, doc))
}

//!-main

//!+visit
// visit appends to links each link found in n and returns the result.
func visit(links []string, n *html.Node) []string {
	//visit函数遍历HTML的节点树，从每一个anchor元素的href属性获得link,将这些links存入字符串数组 中，并返回这个字符串数组。
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	//为了遍历结点n的所有后代结点，每次遇到n的孩子结点时，visit递归的调用自身。这些孩子结点存放在 FirstChild链表中。
	//以下为循环调用
	//for c := n.FirstChild; c != nil; c = c.NextSibling { //赋值给兄弟节点
	//	links = visit(links, c)
	//}
	//以下非循环调用
	if n.FirstChild != nil {
		links = visit(links, n.FirstChild)
	}
	if n.NextSibling != nil {
		links = visit(links, n.NextSibling)
	}
	return links
}

//记录在HTML树中出现的同名元素的次数
func count(res map[string]int, n *html.Node) map[string]int {
	if n.Type == html.ElementNode {
		res[n.Data]++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		res = count(res, c)
	}
	return res
}

//输出所有text结点的内容
func visit3(texts []string, n *html.Node) []string {
	if n.Type == html.TextNode {
		n.Data = strings.Replace(n.Data, " ", "", -1)
		n.Data = strings.Replace(n.Data, "\n", "", -1)
		texts = append(texts, n.Data)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Data == "script" || c.Data == "style" {
			continue
		}

		texts = visit3(texts, c)
	}
	return texts
}

//!-visit

/*
//!+html
package html

type Node struct {
	Type                    NodeType
	Data                    string
	Attr                    []Attribute
	FirstChild, NextSibling *Node
}

type NodeType int32

const (
	ErrorNode NodeType = iota
	TextNode
	DocumentNode
	ElementNode
	CommentNode
	DoctypeNode
)

type Attribute struct {
	Key, Val string
}

func Parse(r io.Reader) (*Node, error)
//!-html
*/
