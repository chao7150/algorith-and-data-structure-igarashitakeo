package main

func swapSliceElems(s []int, x int, y int) {
	tmp := s[x]
	s[x] = s[y]
	s[y] = tmp
}

func bubbleSort(a []int) []int {
	for i := len(a) - 1; 0 < i; i-- {
		for j := 1; j <= i; j++ {
			if a[j-1] > a[j] {
				swapSliceElems(a, j-1, j)
			}
		}
	}
	return a
}

func quickSort(a []int) []int {
	if len(a) <= 1 {
		return a
	}
	var pivot int
	if a[0] < a[1] {
		pivot = a[1]
	} else {
		pivot = a[0]
	}
	former := []int{}
	latter := []int{}
	for _, v := range a {
		if v < pivot {
			former = append(former, v)
		} else {
			latter = append(latter, v)
		}
	}
	return append(quickSort(former), quickSort(latter)...)
}

func MergeSort(a []int) []int {
	if len(a) < 2 {
		return a
	}
	return merge(MergeSort(a[:len(a)/2]), MergeSort((a[len(a)/2:])))
}

func merge(a []int, b []int) []int {
	lena := len(a)
	lenb := len(b)
	itera := 0
	iterb := 0
	res := []int{}
	for {
		if itera < lena && iterb < lenb {
			// aもbも範囲内
			ela := a[itera]
			elb := b[iterb]
			if ela < elb {
				res = append(res, ela)
				itera++
			} else {
				res = append(res, elb)
				iterb++
			}
		} else if itera < lena && lenb == iterb {
			// aは途中だがbは最後まで着いてしまった
			res = append(res, a[itera:]...)
			itera = lena
		} else if lena == itera && iterb < lenb {
			// aは最後まで着いてしまったがbは途中
			res = append(res, b[iterb:]...)
			iterb = lenb
		} else {
			// 両方の要素を取り尽くした
			break
		}
	}
	return res
}

type Heap []int

func (h *Heap) alignUpward(idx int) {
	val := (*h)[idx]
	parentIdx := (idx - 1) / 2
	parentVal := (*h)[parentIdx]
	if val < parentVal {
		(*h)[idx] = parentVal
		(*h)[parentIdx] = val
		h.alignUpward(parentIdx)
	}
}

func (h *Heap) swap(idx1 int, idx2 int) {
	tmp := (*h)[idx1]
	(*h)[idx1] = (*h)[idx2]
	(*h)[idx2] = tmp
}

func (h *Heap) push(i int) {
	*h = append(*h, i)
	h.alignUpward(len(*h) - 1)
}

func (h *Heap) alignDownward(idx int) {
	if len(*h) <= idx*2+1 {
		// 子が存在しない
		// 孫も存在しない
		return

	}
	if len(*h) == idx*2+2 {
		// 子が1つだけ存在する
		if (*h)[idx*2+1] < (*h)[idx] {
			h.swap(idx, idx*2+1)
		}
		// 孫は存在しない
		return
	}
	// 子が2つ存在する
	childIdxA := idx*2 + 1
	childA := (*h)[childIdxA]
	childIdxB := idx*2 + 2
	childB := (*h)[childIdxB]
	if childA <= childB {
		if childA < (*h)[idx] {
			h.swap(idx, childIdxA)
			h.alignDownward(childIdxA)
		}
		return
	}
	if childB < (*h)[idx] {
		h.swap(idx, childIdxB)
		h.alignDownward(childIdxB)
	}
}

func (h *Heap) DeleteMin() int {
	res := (*h)[0]
	(*h)[0] = (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	h.alignDownward(0)
	return res
}

func HeapSort(s []int) []int {
	h := Heap{}
	for _, v := range s {
		h.push(v)
	}
	o := []int{}
	for i := 0; i < len(s); i++ {
		o = append(o, h.DeleteMin())
	}
	return o
}

func (h *Heap) Heapify() {
	for i := (len(*h) - 2) / 2; 0 <= i; i-- {

	}
}

func HeapSortInplace(s []int) {

}

func BucketSort(a []int) []int {
	bucket := make([]int, 10)
	for _, v := range a {
		bucket[v] = v
	}
	out := []int{}
	for _, v := range bucket {
		if v != 0 {
			out = append(out, v)
		}
	}
	return out
}

func DupBucketSort(a []int) []int {
	occurrenceMap := make([]int, 9)
	for _, v := range a {
		occurrenceMap[v]++
	}
	offSet := make([]int, 9)
	for i := range occurrenceMap {
		if i == 0 {
			break
		}
		offSet[i] = offSet[i-1] + occurrenceMap[i-1]
	}
	bucket := make([]int, len(a))
	for _, v := range a {
		bucket[v] = v
	}
	out := []int{}
	for _, v := range bucket {
		if v != 0 {
			out = append(out, v)
		}
	}
	return out
}

type IntAndSortKey struct {
	sortKey int
	value   int
}

func DupBucketSortWithKey(a []IntAndSortKey) []IntAndSortKey {
	occurrenceMap := make([]int, 10)
	for _, v := range a {
		occurrenceMap[v.sortKey]++
	}
	offSet := make([]int, 10)
	for i := range occurrenceMap {
		if i == 0 {
			continue
		}
		offSet[i] = offSet[i-1] + occurrenceMap[i-1]
	}
	bucket := make([]IntAndSortKey, len(a))
	for _, v := range a {
		bucket[offSet[v.sortKey]] = v
		offSet[v.sortKey]++
	}
	out := []IntAndSortKey{}
	for _, v := range bucket {
		if v.sortKey != 0 {
			out = append(out, v)
		}
	}
	return out
}

func PowInt(target int, pow int) int {
	if pow == 0 {
		return 1
	}
	ret := target
	for i := 0; i < pow-1; i++ {
		ret *= ret
	}
	return ret
}

func GetDigit(target int, pos int) int {
	if target < 10 {
		return target
	}
	if pos == 1 {
		return target % 10
	}
	return GetDigit(target/PowInt(10, pos-1), 1)
}

func BucketSortByDigit(a []int, digit int) []int {
	intAndSortKeys := []IntAndSortKey{}
	for _, v := range a {
		sortKey := GetDigit(v, digit)
		intAndSortKeys = append(intAndSortKeys, IntAndSortKey{sortKey, v})
	}
	sorted := DupBucketSortWithKey(intAndSortKeys)
	res := []int{}
	for _, v := range sorted {
		res = append(res, v.value)
	}
	return res
}

func RadixSort4digits(a []int) []int {
	res := a
	for i := 0; i < 4; i++ {
		res = BucketSortByDigit(res, i+1)
	}
	return res
}
