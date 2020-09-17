// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 112.
//!+

// Issues prints a table of GitHub issues matching the search terms.
package main

import (
	"fmt"
	"gopl.io/ch4/github"
	"log"
	"os"
	"time"
)

//!+ go run ch4/issues/main1.go golang   https://api.github.com/search/issues?q=golang
// go run ch4/issues/main.go is:open json decoder
func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	//当前时间，时间戳
	now := time.Now().Unix()
	//前一个月
	preMonth := now - 30*24*3600
	//前一年
	preYear := now - 365*24*3600

	var notMonth []*github.Issue
	var notYear []*github.Issue
	var overYear []*github.Issue
	for _, item := range result.Items {
		createTime := item.CreatedAt.Unix()
		if createTime > preMonth {
			notMonth = append(notMonth, item)
			continue
		}
		if createTime < preMonth && createTime > preYear {
			notYear = append(notYear, item)
			continue
		}
		overYear = append(overYear, item)
	}
	fmt.Println("issues(不到一个月):")
	for _, item := range notMonth {
		fmt.Printf("#%-5d %9.9s %.55s 时间:%s\n", item.Number, item.User.Login, item.Title, item.CreatedAt)
	}

	fmt.Println("issues(不到一年):")
	for _, item := range notYear {
		fmt.Printf("#%-5d %9.9s %.55s 时间:%s\n", item.Number, item.User.Login, item.Title, item.CreatedAt)
	}
	fmt.Println("issues(超过一年):")
	for _, item := range overYear {
		fmt.Printf("#%-5d %9.9s %.55s 时间:%s\n", item.Number, item.User.Login, item.Title, item.CreatedAt)
	}
}

//!-

/*
//!+textoutput
$ go build gopl.io/ch4/issues
$ ./issues repo:golang/go is:open json decoder
13 issues:
#5680    eaigner encoding/json: set key converter on en/decoder
#6050  gopherbot encoding/json: provide tokenizer
#8658  gopherbot encoding/json: use bufio
#8462  kortschak encoding/json: UnmarshalText confuses json.Unmarshal
#5901        rsc encoding/json: allow override type marshaling
#9812  klauspost encoding/json: string tag not symmetric
#7872  extempora encoding/json: Encoder internally buffers full output
#9650    cespare encoding/json: Decoding gives errPhase when unmarshalin
#6716  gopherbot encoding/json: include field name in unmarshal error me
#6901  lukescott encoding/json, encoding/xml: option to treat unknown fi
#6384    joeshaw encoding/json: encode precise floating point integers u
#6647    btracey x/tools/cmd/godoc: display type kind of each named type
#4237  gjemiller encoding/base64: URLEncoding padding is optional
//!-textoutput
*/
