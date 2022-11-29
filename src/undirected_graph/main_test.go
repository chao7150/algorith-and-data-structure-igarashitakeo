package main

import (
	"testing"
)

func TestFill(t *testing.T) {
	res := Fill(5, 2)
	if len(res) != 5 {
		t.Fail()
	}
	for _, v := range res {
		if v != 2 {
			t.Fail()
		}
	}
}

func TestSelectAddingNode(t *testing.T) {
	res := SelectAddingNode([]int{1, 2, 3}, []int{0, 3, 2, 4})
	if res != 2 {
		t.Fail()
	}
}

func TestRemoveFirst(t *testing.T) {
	res := RemoveFirst([]int{1, 2, 3}, 2)
	ans := []int{1, 3}
	for i, v := range res {
		if v != ans[i] {
			t.Fail()
		}
	}
}

func TestCreateSparseAdjacencyMatrix(t *testing.T) {
	res := CreateSparseAdjacencyMatrix(5)
	ans := Matrix{
		{0, Inf, Inf, Inf, Inf},
		{Inf, 0, Inf, Inf, Inf},
		{Inf, Inf, 0, Inf, Inf},
		{Inf, Inf, Inf, 0, Inf},
		{Inf, Inf, Inf, Inf, 0},
	}
	for x, r := range res {
		for y, i := range r {
			if ans[x][y] != i {
				t.Fail()
			}
		}
	}
}

func TestPrim(t *testing.T) {
	m := Matrix{
		{0, 1, Inf, Inf, 8, 3},
		{1, 0, 5, Inf, Inf, 4},
		{Inf, 5, 0, 10, Inf, 6},
		{Inf, Inf, 10, 0, 2, 7},
		{8, Inf, Inf, 2, 0, 9},
		{3, 4, 6, 7, 9, 0},
	}
	ans := Matrix{
		{0, 1, Inf, Inf, Inf, 3},
		{1, 0, 5, Inf, Inf, Inf},
		{Inf, 5, 0, Inf, Inf, Inf},
		{Inf, Inf, Inf, 0, 2, 7},
		{Inf, Inf, Inf, 2, 0, Inf},
		{3, Inf, Inf, 7, Inf, 0},
	}
	res := Prim(m)
	for x, row := range res {
		for y, _ := range row {
			if res[x][y] != ans[x][y] {
				t.Fail()
			}
		}
	}
	m2 := Matrix{
		{0, 5, Inf, Inf, 2},
		{5, 0, 4, 8, 7},
		{Inf, 4, 0, 9, Inf},
		{Inf, 8, 9, 0, 9},
		{2, 7, Inf, 9, 0},
	}
	ans2 := Matrix{
		{0, 5, Inf, Inf, 2},
		{5, 0, 4, 8, Inf},
		{Inf, 4, 0, Inf, Inf},
		{Inf, 8, Inf, 0, Inf},
		{2, Inf, Inf, Inf, 0},
	}
	res = Prim(m2)
	for x, row := range res {
		for y, _ := range row {
			if res[x][y] != ans2[x][y] {
				t.Fail()
			}
		}
	}
}
