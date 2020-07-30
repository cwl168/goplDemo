package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"sort"
	"strings"
	"time"
)

func fetchUrl(url string, ch chan<- map[string]string) { //使用map记录，url -> response
	start := time.Now()
	resp, err := http.Get(url)
	result := make(map[string]string)
	if err != nil {
		result[url] = fmt.Sprintf("http-get: %v", err)
		ch <- result
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	if err != nil {
		result[url] = fmt.Sprintf("while reading:%v %v", url, err)
		ch <- result
		return
	}

	resp.Body.Close()
	secs := time.Since(start).Seconds()
	//fmt.Printf("%v %v %v Bytes %v's\n",url,resp.Status,nbytes,secs)
	result[url] = fmt.Sprintf("%v %v %v Bytes %v's", url, resp.Status, nbytes, secs)
	ch <- result

}

//go run ch1/fetchall/main2.go baidu.com sina.cn 163.com google.cn qq.com weibo.cn
func main() {
	start := time.Now()
	ch := make(chan map[string]string)
	count := 0
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http") {
			url = "https://" + url
		}
		fmt.Printf("start fetch %v\n", url)
		go fetchUrl(url, ch) // 重复3次
		go fetchUrl(url, ch)
		go fetchUrl(url, ch)
		count += 3
	}
	result := make(map[string]string)
	keys := []string{}
	fmt.Println(count)
	//为什么要循环count
	//for i := 0; i < count; i++ {
	for k, v := range <-ch {
		result[k] = v
		keys = append(keys, k)
	}
	//}
	sort.Strings(keys) //按照url排序
	for _, k := range keys {
		fmt.Printf("%s : %s\n", k, result[k]) //输出按照URL排序后的结果
	}
	fmt.Printf("done, use %v seconds", time.Since(start).Seconds())
}
