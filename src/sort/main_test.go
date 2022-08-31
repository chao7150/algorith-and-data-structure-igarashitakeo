package main

import (
	"testing"
)

func TestSort(t *testing.T) {
	a := []int{8, 4, 5, 7, 6, 1, 2, 9, 3}
	b := MergeSort(a)
	c := 0
	for i, v := range b {
		if v < c {
			t.Fatalf(`%v th value %v is smaller than previous value %v`, i, v, c)
		}
		c = v
	}
}

func TestMerge(t *testing.T) {
	a := []int{4, 5, 7, 8}
	b := []int{1, 2, 3, 6, 9}
	c := merge(a, b)
	for i, v := range c {
		if v != i+1 {
			t.Fail()
		}
	}
}
