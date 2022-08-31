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

func TestAlign(t *testing.T) {
	h := Heap{1, 3, 6, 4, 8, 7, 10, 5, 9, 11, 2}
	h.align(10)
	res := []int{1, 2, 6, 4, 3, 7, 10, 5, 9, 11, 8}
	for i, v := range h {
		if v != res[i] {
			t.Fail()
		}
	}
}

func TestPush(t *testing.T) {
	h := Heap{}
	h.push(31)
	h.push(21)
	h.push(48)
	h.push(9)
	res := []int{9, 21, 48, 31}
	for i, v := range h {
		if v != res[i] {
			t.Fail()
		}
	}
}
