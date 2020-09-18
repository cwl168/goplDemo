package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
	"strings"
)

//go run ch5/findlinks2/main1.go  https://www.baidu.com
func main() {
	words, images, _ := CountWordsAndImages(os.Args[1])
	fmt.Printf("文字：%d,图片：%d \n", words, images)
}

// CountWordsAndImages does an HTTP GET request for the HTML
// document url and returns the number of words and images in it.
func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(os.Stdin)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	//bare return
	return
}
func countWordsAndImages(n *html.Node) (words, images int) {

	texts, images := visit3(nil, 0, n)
	for _, v := range texts {
		v = strings.Trim(strings.TrimSpace(v), "\r\n")
		if v == "" {
			continue
		}
		words += strings.Count(v, "")
	}
	//bare return
	return
}

//递归循环html
func visit3(texts []string, imgs int, n *html.Node) ([]string, int) {
	//文本
	if n.Type == html.TextNode {
		texts = append(texts, n.Data)
	}
	//图片
	if n.Type == html.ElementNode && (n.Data == "img") {
		imgs++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Data == "script" || c.Data == "style" {
			continue
		}

		texts, imgs = visit3(texts, imgs, c)
	}
	//多返回值
	return texts, imgs
}
