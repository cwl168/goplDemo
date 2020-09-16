// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 101.

// Package treesort provides insertion sort using an unbalanced binary tree.
package treesort

import (
	"fmt"
	"os"
)

//!+
type tree struct {
	value       int
	left, right *tree
}

// Sort sorts values in place.
func Sort(values []int) {
	var root *tree
	fmt.Println(values)

	for _, v := range values {
		root = add(root, v) //二叉排序树，中序遍历有序
	}
	fmt.Println("前序遍历")
	trans(root)
	fmt.Println()
	fmt.Println("中序遍历")
	trans1(root)
	fmt.Println()
	fmt.Println("后序遍历")
	trans2(root)
	fmt.Println(values[:0])
	fmt.Println(appendValues(values[:0], root))
	os.Exit(0)
}

//中序遍历的结果加入到slice中
// appendValues appends the elements of t to values in order
// and returns the resulting slice.
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		// Equivalent to return &tree{value: value}.
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}
func trans(root *tree) {
	//前序遍历：先遍历根节点，再遍历左子树，再遍历右子树
	if root == nil {
		return
	}
	fmt.Printf("%d ", root.value)
	trans(root.left)
	trans(root.right)
}
func trans1(root *tree) {
	//中序遍历:先遍历左子树，再遍历根节点，再遍历右子树
	if root == nil {
		return
	}
	trans1(root.left)
	fmt.Printf("%d ", root.value)
	trans1(root.right)
}
func trans2(root *tree) {
	//后序遍历:先遍历左子树，再遍历右子树，再遍历根节点
	if root == nil {
		return
	}
	trans2(root.left)
	trans2(root.right)
	fmt.Printf("%d ", root.value)
}

//!-
