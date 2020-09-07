//非递归实现
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(compare("abec", "ecababc"))
}
func compare(str1 string, str2 string) bool {
	//多层循环的遍历，我们的一个原则是内大外小
	num1 := strings.Count(str1, "")
	num2 := strings.Count(str2, "")
	if num2 < num1 {
		str1, str2 = str2, str1
	}
	var res bool
	count := 0 //循环次数
	for _, v := range str1 {
		res = false
		for _, sv := range str2 {
			count++
			if v == sv {
				res = true
			}
		}
		if !res {
			break
		}
	}
	fmt.Println(count)
	return res
}
