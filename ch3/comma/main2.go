//非递归实现
package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}
func comma(s string) string {
	var newByte byte = ','
	n := len(s)
	buf := bytes.NewBuffer([]byte{}) //创建bytes.buffer
	if n <= 3 {
		return s
	}
	for i := 0; i < n; i++ {
		if (n-i)%3 == 0 && i != 0 {
			buf.WriteByte(newByte)
		}
		buf.WriteByte(s[i])
	}
	return buf.String()
}

/*
i n-i
0 8   1
1 7	  2
2 6   , 3
3 5   4
4 4   5
5 3   , 6
6 2   7
7 1   8
*/
