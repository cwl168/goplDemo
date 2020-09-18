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

//go run ch5/outline2/outline3.go https://www.baidu.com
//练习 5.8： 修改pre和post函数，使其返回布尔类型的返回值。返回false时，中止forEachNoded的遍 历。使用修改后的代码编写ElementByID函数，根据用户输入的id查找第一个拥有该id元素的HTML元素， 查找成功后，停止遍历。
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

	//!+call
	fmt.Println(ElementByID(doc, "noscript", startElement2))
	//!-call

	return nil
}

func ElementByID(n *html.Node, idStr string, pre func(*html.Node, string) bool) *html.Node {
	//显式的调用一下
	if pre != nil {
		if pre(n, idStr) {
			return n
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ElementByID(c, idStr, pre)
	}
	return n
}

func startElement2(n *html.Node, idStr string) bool {
	if n.Type == html.ElementNode {
		for _, a := range n.Attr {
			if a.Key == "id" && a.Val == idStr {
				fmt.Println(a.Val)
				return true
				break

			}
		}
	}
	return false
}
