package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

//在http包使用的时候，注册了/这个根路径的模式处理，浏览器会自动的请求favicon.ico
// ab -n 10000 -c 100 -s 60 -k 127.0.0.1:8000/
func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
	log.Printf("Count %d\n", count)
}

// handler echoes the Path component of the requested URL.
func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.RequestURI() != "/favicon.ico" {
		mu.Lock()
		count++
		mu.Unlock()
		fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
		log.Printf("URL.Path = %q\n", r.URL.Path)
	}

}

// counter echoes the number of calls so far.
func counter(w http.ResponseWriter, r *http.Request) {
	if r.URL.RequestURI() != "/favicon.ico" {
		mu.Lock()
		fmt.Fprintf(w, "Count %d\n", count)
		mu.Unlock()
		log.Printf("Count %d\n", count)
	}

}
