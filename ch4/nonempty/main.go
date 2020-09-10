// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 91.

//!+nonempty

// Nonempty is an example of an in-place slice algorithm.
package main

import "fmt"

// nonempty returns a slice holding only the non-empty strings.
// The underlying array is modified during the call.
//nonempty函数将在原有slice内存空间之上返回不包含空字符串的列表
func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	fmt.Println(strings, i)
	return strings[:i]
}

//!-nonempty

func main() {
	//!+main
	/*data := []string{"one", "", "three"}
	fmt.Printf("%q\n", data)
	fmt.Printf("%q\n", nonempty(data)) // `["one" "three"]`		%q 使用Go语法以及必须时使用转义，以双引号括起来的字符串或者字节切片[]byte，或者是以单引号括起来的数字
	fmt.Printf("%q\n", data)           // `["one" "three" "three"]`   输入的slice和输出的slice共享一个底层数组。这可以避免分配另一个数组，不过 原来的数据将可能会被覆盖*/

	data := []string{"one", "", "three"}
	data1 := nonempty2(data)
	fmt.Printf("%q\n", data)
	data1 = append(data1, "four")
	fmt.Printf("%q\n", data1)

	s := []int{5, 6, 7, 8, 9}
	fmt.Println(remove(s, 2))
	//!-main
}

//!+alt
func nonempty2(strings []string) []string {
	out := strings[:0] // zero-length slice of original
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}
func remove(slice []int, i int) []int {
	fmt.Println(slice[i:])
	fmt.Println(slice[i+1:])
	copy(slice[i:], slice[i+1:]) //slice[i:] 7 8 9   slice[i+1:]  8 9   结果 [5 6 8 9 9]
	fmt.Println(slice)
	return slice[:len(slice)-1] //[5 6 8 9]
}

//!-alt
