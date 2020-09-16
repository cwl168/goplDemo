// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package treesort_test

import (
	"sort"
	"testing"

	"gopl.io/ch4/treesort"
)

func TestSort(t *testing.T) {
	//data := make([]int, 10)
	//for i := range data {
	//	data[i] = rand.Int() % 10
	//}
	data := []int{2, 10, 8, 5, 30, 4, 7, 9}
	treesort.Sort(data)
	if !sort.IntsAreSorted(data) {
		t.Errorf("not sorted: %v", data)
	}
}
