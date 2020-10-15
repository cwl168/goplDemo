// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 223.

// Netcat is a simple read/write client for TCP servers.
package main

import (
	"io"
	"log"
	"net"
	"os"
)

//!+
//配合 ch8/reverb1/reverb.go服务端运行
//当main goroutine从标准输入流中读取内容并将其发送给服务器时，另一个goroutine会读取并打印服务 端的响应。
//第三次shout在前一个shout处理完成之前一直没有被处理，
func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	//，它在主goroutine中（译注：就是执行main函数的goroutine）将标准输入复制到 server，因此当客户端程序关闭标准输入时，后台goroutine可能依然在工作。
	go mustCopy(os.Stdout, conn) //另一个goroutine会读取并打印服务 端的响应。
	mustCopy(conn, os.Stdin)     //main goroutine从标准输入流中读取内容并将其发送给服务器
}

//!-

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

/**
➜  gopl.io git:(master) ✗  go run ch8/netcat2/netcat.go
Hello?
         HELLO?
         Hello?
         hello?
Is there anybody there?
         IS THERE ANYBODY THERE?
         Is there anybody there?
         is there anybody there?
Is there anybody there?
         IS THERE ANYBODY THERE?
erter
         Is there anybody there?
         is there anybody there?
         ERTER
         erter
         erter
aaaa
         AAAA
bbb
         aaaa
         aaaa
         BBB
         bbb
         bbb

*/
