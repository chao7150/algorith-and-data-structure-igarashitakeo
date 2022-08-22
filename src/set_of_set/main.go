package main

type Item struct {
	isRoot bool
	value  int64
	parent *Item
}

func find(i *Item) int64 {
	if i.isRoot {
		return i.value
	}
	return find(i.parent)
}
