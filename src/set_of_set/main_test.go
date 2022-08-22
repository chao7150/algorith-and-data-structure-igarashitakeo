package main

import (
	"testing"
)

func TestFind(t *testing.T) {
	ir := Item{true, 0, nil}
	i1 := Item{false, 1, &ir}
	i2 := Item{false, 2, &i1}
	r := find(&i2)
	if r != 0 {
		t.Fail()
	}
}

func createSingleNode(value int64) *Item {
	root := Item{true, value, nil}
	node := Item{false, value, &root}
	return &node
}

func TestMerge(t *testing.T) {
	
}
