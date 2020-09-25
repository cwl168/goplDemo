// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 165.

// Package intset provides a set of integers based on a bit vector.
package main

import (
	"bytes"
	"fmt"
)

//!+intset

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint64 //无符号64位整数,8字节
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)                             // x/64的商作为字的下标， x%64得到的值作为这个字内的bit的所在位置。
	return word < len(s.words) && s.words[word]&(1<<bit) != 0 //将1左移bit后，那个位置自然就是1，然后和以前的数据做&(位与,两位同时为“1”，结果才为“1”，否则为0)，判断是否为0即可
}
func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)
	//将1左移bit后，那个位置自然就是1，然后对取反，再与当前值做&，即可清除当前的位置了
	s.words[word] ^= (1 << bit) //左移 异或
	//s.words[word] &= ^ (1 << bit)  // 左移 取反 位与
	//s.words[word] &= s.words[word] ^ (1 << bit)   //左移 异或 位与
	//以上三种都可以实现删除某一个元素
}

// Add adds the non-negative value x to the set.

func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit //将1左移bit后，那个位置自然就是1，然后和以前的数据做|，这样，那个位置就替换成1了
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

//清空
func (s *IntSet) Clear() {
	for i, word := range s.words {
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				s.words[i] ^= 1 << uint(j) //左移 异或
			}
		}
	}
}
func (s *IntSet) Len() int {
	sum := 0
	for _, word := range s.words {
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				sum++
			}
		}
	}
	return sum
}

//!-intset

//!+string

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 { //位于 两位同时为“1”，结果才为“1”，否则为0
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}
func (s *IntSet) AddAll(vals ...int) (int, error) {
	if len(vals) == 0 {
		return 0, fmt.Errorf("至少传一个参数")
	}
	var sum int
	for _, val := range vals {
		sum += val
	}
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				sum += 64*i + j
			}
		}
	}
	return sum, nil
}

//!-string
