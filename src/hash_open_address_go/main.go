package hashopenaddressgo

type IntValue struct {
	value   int
	set     bool
	deleted bool
}

type Hash [10]IntValue

func getIndex(value int) int {
	return value % 10
}

func (h *Hash) member(value int) bool {
	var count int
	for {
		i := getIndex((value) + count)
		if h[i].set && h[i].value == value {
			return true
		}
		if !h[i].set {
			return false
		}
		count++
	}
}

func (h *Hash) insert(value int) {
	if h.member(value) {
		return
	}
	count := 0
	for {
		i := getIndex(value) + count
		if !h[i].set || h[i].deleted {
			h[i] = IntValue{value: value, set: true}
			return
		}
		count++
	}
}

func (h *Hash) delete(value int) {
	c := 0
	for {
		i := getIndex(value) + c
		if h[i].set && h[i].value == value {
			h[i].deleted = true
			h[i].set = false
			return
		}
		c++
	}
}
