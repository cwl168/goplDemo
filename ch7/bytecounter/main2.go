package main

import (
	"fmt"
	"io"
	"os"
)

type CountWriter struct {
	Writer io.Writer
	Count  int
}

func (cw *CountWriter) Write(content []byte) (int, error) { //CountWriter 结构体实现了接口 io.Writer
	n, err := cw.Writer.Write(content) //os/file.go  Write方法
	if err != nil {
		return n, err
	}
	cw.Count += n
	return n, nil
}

func CountingWriter(writer io.Writer) (io.Writer, *int) {
	cw := CountWriter{
		Writer: writer,
	}
	return &cw, &(cw.Count)
}

// 写一个带有如下函数签名的函数CountingWriter，传入一个io.Writer接口类型，返回一个新的Writer类型把原来的Writer封装在里面和一个表示写入新的Writer字节数的int64类型指针
func main() {
	cw, counter := CountingWriter(os.Stdout)                  // os.Stdout 为 io.Writer 类型
	fmt.Fprintf(cw, "%s", "Print somethind to the screen...") //长度为32
	fmt.Println(*counter)
	cw.Write([]byte("Append soething...")) //追加 长度为50
	fmt.Println(*counter)
}
