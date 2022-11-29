package main

import "testing"

func TestKMP(t *testing.T) {
	text := "にわにはにわにわとりがいる"
	pattern := "にわとり"
	slideMap := []int{1, 1, 1, 1}
	res := KMP(text, pattern, slideMap)
	if res != 6 {
		t.Fail()
	}
	res = KMP("ながの", "たい", slideMap)
	if res != 0 {
		t.Fail()
	}
	res = KMP("ながの", "のい", slideMap)
	if res != 0 {
		t.Fail()
	}
	res = KMP("a", "abc", slideMap)
	if res != 0 {
		t.Fail()
	}
}
