package main

import (
	"flag"
	"io"
	"log"
	"net"
	"time"
)

//支持传入参数作为端口号
var port = flag.String("port", "8000", "请输入端口")

//go run ch8/clock2/clock2.go -port 8000
//go run ch8/clock2/clock2.go -port 8001
//go run ch8/clock2/clock2.go -port 8002
func main() {
	flag.Parse()
	listener, err := net.Listen("tcp", "localhost:"+*port)
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn) //新建goroutines处理连接
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // e.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}
