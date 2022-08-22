package hashopenaddressgo

import (
	"testing"
)

func TestHash(t *testing.T) {
	h := getIndex(0)
	if getIndex(0) != 0 {
		t.Fatalf(`hash(0) = %d`, h)
	}
}

func TestMember(t *testing.T) {
	h := Hash{}
	h.insert(5)
	if !h.member(5) {
		t.Fatalf("should be true")
	}
	if h.member(6) {
		t.Fatalf("should be false")
	}
}

func TestInsert(t *testing.T) {
	h := Hash{}
	h.insert(5)
	if !h.member(5) {
		t.Fail()
	}
	h.insert(15)
	if !h.member(5) || !h.member(15) {
		t.Fail()
	}
	h.insert(5)
	h.delete(5)
	if h.member(5) {
		t.Fail()
	}
}

func TestDelete(t *testing.T) {
	h := Hash{}
	h.insert(5)
	h.insert(15)
	h.delete(5)
	if h.member(5) {
		t.Fail()
	}
	h.delete(15)
	if h.member(15) {
		t.Fail()
	}
}
