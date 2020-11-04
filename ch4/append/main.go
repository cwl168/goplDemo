// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 88.

// Append illustrates the behavior of the built-in append function.
package main

import (
	"fmt"
	"os"
)

func appendslice(x []int, y ...int) []int {
	var z []int
	zlen := len(x) + len(y)
	if zlen <= cap(x) {
		// There is room to expand the slice.
		z = x[:zlen]
	} else {
		// There is insufficient space.
		// Grow by doubling, for amortized linear complexity.
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	copy(z[len(x):], y)
	return z
}

//!+append
func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1  // 1 2,3,4....
	if zlen <= cap(x) { //每次调用appendInt函数，必须先检测slice底层数组是否有足够的容量来保存新添加的元素
		// There is room to grow.  Extend the slice. 还有增长空间。扩展部分(存在空间)
		z = x[:zlen] //获取索引 0 到 zlen-1的元素
		//fmt.Println(x,zlen,z,len(x))
	} else {
		// There is insufficient space.  Allocate a new array.没有足够的空间。分配一个新数组。
		// Grow by doubling, for amortized linear complexity. 通过加倍增长，平摊线性复杂度
		/**
		len(x) + 1 < 2 * len(x) 两边 -len(x)
		就变成了 1 < len（x） 就是切片长度大于1时，就会扩容成两倍，小于1时，即扩容1个元素
		*/
		zcap := zlen
		if 1 < len(x) { // zcap < 2 * len(x)   i=2,4,8  就会扩容2倍
			zcap = 2 * len(x)
		}
		fmt.Println(zlen, zcap, len(x))
		z = make([]int, zlen, zcap)
		copy(z, x) // a built-in function; see text
	}
	//z切片中添加新数据 y
	z[len(x)] = y
	return z
}

//!-append

//!+growth
func main() {
	//array := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	//fmt.Println(array[:3])
	var x, y []int
	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		fmt.Printf("%d  cap=%d\t%v\t%v\n", i, cap(y), y, x)
		x = y
	}
	fmt.Println()
	os.Exit(0)
	for i := 0; i < 10; i++ {
		y = appendslice(x, i)
		fmt.Printf("%d  cap=%d\t%v\t%v\n", i, cap(y), y, x)
		x = y
	}
}

//!-growth

/*
//!+output
0  cap=1   [0]
1  cap=2   [0 1]
2  cap=4   [0 1 2]
3  cap=4   [0 1 2 3]
4  cap=8   [0 1 2 3 4]
5  cap=8   [0 1 2 3 4 5]
6  cap=8   [0 1 2 3 4 5 6]
7  cap=8   [0 1 2 3 4 5 6 7]
8  cap=16  [0 1 2 3 4 5 6 7 8]
9  cap=16  [0 1 2 3 4 5 6 7 8 9]
//!-output
*/
