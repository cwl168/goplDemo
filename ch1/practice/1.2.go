package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	data := strings.Join(os.Args[1:], " ") //字符串链接，把slice a通过sep链接起来
	fmt.Println(data)
	fmt.Println(len(data))
	for k, v := range os.Args[1:] {
		fmt.Println(k, v)
	}
}
