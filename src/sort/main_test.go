package main

import (
	"testing"
)

func TestSort(t *testing.T) {
	a := []int{8, 4, 5, 7, 6, 1, 2, 9, 3}
	b := HeapSort(a)
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

func TestAlignUpward(t *testing.T) {
	h := Heap{1, 3, 6, 4, 8, 7, 10, 5, 9, 11, 2}
	h.alignUpward(10)
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

func TestHeapSwap(t *testing.T) {
	h := Heap{1, 2}
	h.swap(0, 1)
	res := []int{2, 1}
	AssertSliceEqual(t, h, res)
}

func AssertSliceEqual(t *testing.T, s []int, res []int) {
	for i, v := range s {
		if v != res[i] {
			t.Fatalf(`expected: %v, given: %v`, res, s)
		}
	}
}

func TestHeapAlignDownward(t *testing.T) {
	h := Heap{0}
	h.alignDownward(0)
	res := []int{0}
	AssertSliceEqual(t, h, res)

	h = Heap{1, 0}
	h.alignDownward(0)
	res = []int{0, 1}
	AssertSliceEqual(t, h, res)

	h = Heap{0, 1}
	h.alignDownward(0)
	res = []int{0, 1}
	AssertSliceEqual(t, h, res)

	h = Heap{2, 1, 0}
	h.alignDownward(0)
	res = []int{0, 1, 2}
	AssertSliceEqual(t, h, res)

	h = Heap{3, 2, 4, 0}
	h.alignDownward(0)
	res = []int{2, 0, 4, 3}
	AssertSliceEqual(t, h, res)

	h = Heap{9, 3, 6, 6, 8, 10, 12, 10, 13, 11}
	h.alignDownward(0)
	res = []int{3, 6, 6, 9, 8, 10, 12, 10, 13, 11}
	AssertSliceEqual(t, h, res)
}

func TestHeapDeleteMin(t *testing.T) {
	h := Heap{1, 3, 6, 6, 8, 10, 12, 10, 13, 11, 9}
	min := h.DeleteMin()
	res := []int{3, 6, 6, 9, 8, 10, 12, 10, 13, 11}
	AssertSliceEqual(t, h, res)
	if min != 1 {
		t.Fail()
	}
	min = h.DeleteMin()
	res = []int{6, 8, 6, 9, 11, 10, 12, 10, 13}
	AssertSliceEqual(t, h, res)
	if min != 3 {
		t.Fail()
	}
}

func TestBucketSort(t *testing.T) {
	a := []int{8, 4, 5, 7, 6, 1, 2, 9, 3}
	b := BucketSort(a)
	c := 0
	for i, v := range b {
		if v < c {
			t.Fatalf(`%v th value %v is smaller than previous value %v`, i, v, c)
		}
		c = v
	}
}

func TestDupBucketSort(t *testing.T) {
	a := []int{3, 2, 2, 1, 2, 2, 1, 3, 3, 1, 2, 3}
	b := DupBucketSort(a)
	c := 0
	for i, v := range b {
		if v < c {
			t.Fatalf(`%v th value %v is smaller than previous value %v`, i, v, c)
		}
		c = v
	}
}

func TestDupBucketSortWithKey(t *testing.T) {
	a := []IntAndSortKey{{4, 24}, {4, 14}, {1, 11}}
	b := DupBucketSortWithKey(a)
	res := []int{11, 24, 14}
	for i, v := range b {
		if res[i] != v.value {
			t.Fail()
		}
	}
	a = []IntAndSortKey{{1, 251}, {1, 151}, {4, 334}, {4, 114}, {4, 114}, {3, 513}, {5, 365}, {3, 363}, {1, 1}, {1, 1001}, {2, 42}}
	b = DupBucketSortWithKey(a)
	res = []int{251, 151, 1, 1001, 42, 513, 363, 334, 114, 114, 365}
	for i, v := range b {
		if res[i] != v.value {
			t.Fail()
		}
	}
}

func TestPowInt(t *testing.T) {
	if PowInt(10, 0) != 1 {
		t.Fail()
	}
	if PowInt(10, 1) != 10 {
		t.Fail()
	}
	if PowInt(10, 2) != 100 {
		t.Fail()
	}
}

func TestGetDigit(t *testing.T) {
	res := GetDigit(12345, 1)
	if res != 5 {
		t.Fail()
	}
	res = GetDigit(12345, 2)
	if res != 4 {
		t.Fail()
	}
	res = GetDigit(12345, 3)
	if res != 3 {
		t.Fail()
	}
	res = GetDigit(12345, 6)
	if res != 0 {
		t.Fail()
	}
}

func TestRadixSort(t *testing.T) {
	a := []int{251, 151, 334, 114, 114, 513, 365, 363, 1, 1001, 42}
	b := RadixSort4digits(a)
	c := 0
	for i, v := range b {
		if v < c {
			t.Fatalf(`%v th value %v is smaller than previous value %v`, i, v, c)
		}
		c = v
	}
}
