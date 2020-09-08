// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 83.

// The sha256 command computes the SHA256 hash (an array) of a string.
package main

//!+
import (
	"crypto/sha256"
	"fmt"
)

func main() {
	//crypto/sha256包的Sum256函数对一个任意的字节slice类型的数据生成一个对应 的消息摘要。

	fmt.Println(compareSha256("x", "X"))
	// Output:
	// 2d711642b726b04401627ca9fbac32f5c8530fb1903cc4db02258717921a4881   每2位一个代表 ASCII字符，例如 2d位45
	// 4b68ab3847feda7d6c62c1fbcbeebfa35eab7351ed5e78f4ddadea5df64b8015
	// false
	// [32]uint8
}

//编写一个函数，计算两个SHA256哈希码中不同bit的数目
//[00101101 01110001 00010110 01000010 10110111 00100110 10110000 01000100 00000001 01100010 01111100 10101001 11111011 10101100 00110010 11110101 11001000 01010011 00001111 10110001 10010000 00111100 11000100 11011011 00000010 00100101 10000111 00010111 10010010 00011010 01001000 10000001]
//[01001011 01101000 10101011 00111000 01000111 11111110 11011010 01111101 01101100 01100010 11000001 11111011 11001011 11101110 10111111 10100011 01011110 10101011 01110011 01010001 11101101 01011110 01111000 11110100 11011101 10101101 11101010 01011101 11110110 01001011 10000000 00010101]
func compareSha256(str1 string, str2 string) int {
	a := sha256.Sum256([]byte(str1))
	b := sha256.Sum256([]byte(str2))
	num := 0
	//循环字节数组
	for i := 0; i < len(a); i++ {
		t := a[i] ^ b[i] //两个位相同为0，相异为1
		for t != 0 {
			t &= (t - 1) //两个位都为1时，结果才为1
			num++
		}
	}
	return num
}

//!-
