// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// +build ignore

// The thumbnail command produces thumbnails of JPEG files
// whose names are provided on each line of the standard input.
//
// The "+build ignore" tag (see p.295) excludes this file from the
// thumbnail package, but it can be compiled as a command and run like
// this:
//
// Run with:
//   $ go run $GOPATH/src/gopl.io/ch8/thumbnail/main.go
//   foo.jpeg
//   ^D
//
package main

import (
	"fmt"
	"gopl.io/ch8/thumbnail"
)

func main() {
	filenames:= []string{"/Users/caoweilin/go/src/gopl.io/ch8/thumbnail/img/1.jpg","/Users/caoweilin/go/src/gopl.io/ch8/thumbnail/img/2.jpg"}
	ch := make(chan struct{})
	for _, f := range filenames {
		fmt.Println(f)
		go func(f string) {
			thumbnail.ImageFile(f) // NOTE: ignoring errors
			ch <- struct{}{}
		}(f)  //必须加参数，否则会出现循环变量快照问题
	}

	// Wait for goroutines to complete.
	for range filenames {
		fmt.Println(<-ch)
	}


	/*input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		thumb, err := thumbnail.ImageFile(input.Text())
		if err != nil {
			log.Print(err)
			continue
		}
		fmt.Println(thumb)
	}
	if err := input.Err(); err != nil {
		log.Fatal(err)
	}*/
}
