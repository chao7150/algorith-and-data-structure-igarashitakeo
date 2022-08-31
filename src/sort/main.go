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

func (h *Heap) align(idx int) {
	val := (*h)[idx]
	parentIdx := (idx - 1) / 2
	parentVal := (*h)[parentIdx]
	if val < parentVal {
		(*h)[idx] = parentVal
		(*h)[parentIdx] = val
		h.align(parentIdx)
	}
}

func (h *Heap) push(i int) {
	*h = append(*h, i)
	h.align(len(*h) - 1)
}
