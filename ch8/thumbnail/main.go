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
	"log"
	"os"
	"sync"

	"gopl.io/ch8/thumbnail"
)

func main() {
	//测试makeThumbnails2
	/*filenames:= []string{"/Users/caoweilin/go/src/gopl.io/ch8/thumbnail/img/1.jpg","/Users/caoweilin/go/src/gopl.io/ch8/thumbnail/img/2.jpg"}
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
	}*/

	//测试makeThumbnails6    sizes无缓冲channel
	/*filenames:= []string{"/Users/caoweilin/go/src/gopl.io/ch8/thumbnail/img/1.jpg","/Users/caoweilin/go/src/gopl.io/ch8/thumbnail/img/2.jpg"}
	sizes := make(chan int64)
	var wg sync.WaitGroup // number of working goroutines
	for _,f := range filenames {
		wg.Add(1)
		// worker
		go func(f string) {
			defer wg.Done()
			thumb, err := thumbnail.ImageFile(f)
			if err != nil {
				log.Println(err)
				return
			}
			info, _ := os.Stat(thumb) // OK to ignore error
			sizes <- info.Size()
		}(f)
	}

	//在关闭掉sizes channel 之前work们退出
	// closer
	go func() {
		wg.Wait()
		close(sizes)
	}()

	//在main goroutine中使用了range loop来计 算总和
	var total int64
	for size := range sizes {
		total += size
	}
	fmt.Println(total)*/


	//测试makeThumbnails6    sizes有缓冲channel
	filenames:= []string{"/Users/caoweilin/go/src/gopl.io/ch8/thumbnail/img/1.jpg","/Users/caoweilin/go/src/gopl.io/ch8/thumbnail/img/2.jpg"}
	sizes := make(chan int64,2)
	var wg sync.WaitGroup // number of working goroutines
	for _,f := range filenames {
		wg.Add(1)
		// worker
		go func(f string) {
			defer wg.Done()
			thumb, err := thumbnail.ImageFile(f)
			if err != nil {
				log.Println(err)
				return
			}
			info, _ := os.Stat(thumb) // OK to ignore error
			sizes <- info.Size()
		}(f)
	}
	// closer
	wg.Wait()
	close(sizes)
	//在main goroutine中使用了range loop来计 算总和
	var total int64
	for size := range sizes {
		total += size
	}
	fmt.Println(total)








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
