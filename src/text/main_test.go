package main

import "testing"

func TestKMP(t *testing.T) {
	text := "にわにはにわにわとりがいる"
	pattern := "にわとり"
	res := KMP(text, pattern, []int{1, 1, 1, 1})
	if res != 6 {
		t.Fail()
	}
	res = KMP("ながの", "たい", []int{1, 1})
	if res != 0 {
		t.Fail()
	}
	res = KMP("ながの", "のい", []int{1, 1})
	if res != 0 {
		t.Fail()
	}
	res = KMP("a", "abc", []int{1, 1, 1})
	if res != 0 {
		t.Fail()
	}
}
