package main

import (
	"fmt"
	"io"
)

type StringReader struct {
	data    string
	current int
}

func (sr *StringReader) Read(b []byte) (n int, err error) {
	if len(b) == 0 { // 不需要读入
		return 0, nil
	}
	// copy() guarantee copy min(len(b),len(sr.data[sr.current:])) bytes
	n = copy(b, sr.data[sr.current:])
	if sr.current += n; sr.current >= len(sr.data) { // 已读完
		err = io.EOF
	}

	return
}

func NewReader(in string) *StringReader {
	sr := new(StringReader)
	sr.data = in
	return sr
}

// strings.NewReader函数通过读取一个string参数返回一个满足io.Reader接口类型的值（和其它值）。实现一个简单版本的NewReader，并用它来构造一个接收字符串输入的HTML解析器（§5.2）
// 主要判断b长度为0，且考虑读完返回EOF，修改测试读取逻辑
func main() {
	str := "Hello World"
	sr := NewReader(str)
	data := make([]byte, 10)    // 每次最多读10个byte
	n, err := sr.Read(data[:0]) // 初始化
	fmt.Println(data[:0], n, err)
	for err == nil {
		n, err = sr.Read(data)
		fmt.Println(n, string(data[0:n]), err) // 重新取切片，因为最后一次结果data[n:]含有上一轮的结果
	}
	//output:
	// 10 Hello Worl
	// 1 d
}
