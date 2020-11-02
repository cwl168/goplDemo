// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 10.
//!+

// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

//1.   go run ch1/dup2/main.go ch1/dup2/test.txt  ch1/dup2/test2.txt
//2.   go run ch1/dup2/main.go scanner对象从程序的标准输入中读取内容,Ctrl + d结束

//Exercise 1.4: Modify dup2 to print the names of all files in which each duplicated line occurs.
/**
map[
12:[ch1/dup2/test.txt]
23:[ch1/dup2/test.txt ch1/dup2/test.txt]
34:[ch1/dup2/test.txt] aaa:[ch1/dup2/test2.txt]
bb:[ch1/dup2/test2.txt ch1/dup2/test2.txt ch1/dup2/test2.txt]
c:[ch1/dup2/test2.txt]
ccc:[ch1/dup2/test2.txt]
]
*/
func main() {
	counts := make(map[string][]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	fmt.Println(counts)
	for line, files := range counts {
		if len(files) > 1 {
			fmt.Printf("%d\t%s\t%s\n", len(files), line, files)
		}
	}
}

func countLines(f *os.File, counts map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()] = append(counts[input.Text()], f.Name())
	}
}

//!-
