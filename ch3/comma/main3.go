//非递归实现
package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	fmt.Println(comma(-5123456.23))
}
func comma(str float64) string {
	//整型转换成字符串
	s := fmt.Sprintf("%.2f", str)
	//取出小数点后面部分
	var end string
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		end = s[dot:] //.23小数部分
		s = s[:dot]   // 整数部分
	}
	fmt.Println(end)
	num := len(s)
	var buf bytes.Buffer
	j := 1
	for i := num - 1; i >= 0; i-- {
		buf.WriteByte(s[i])
		if j%3 == 0 && i != 0 {
			buf.WriteString(",")
		}
		j++
	}
	res := buf.String()
	var r bytes.Buffer
	//反转字符串
	for i := len(res) - 1; i >= 0; i-- {
		r.WriteByte(res[i])
	}
	r.WriteString(end)
	return r.String()
}
