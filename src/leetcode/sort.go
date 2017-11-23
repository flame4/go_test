package leetcode

import (
	"fmt"
)

func quick_sort(l []int, head int, tail int) {
	if head > tail {
		return
	}
	fmt.Println(l)
	mid := sort(l, head, tail)
	quick_sort(l, head, mid-1)
	quick_sort(l, mid+1, tail)
}

func sort(l []int, head int, tail int) int {
	for head < tail {
		for head < tail && l[head] < l[tail] {
			tail--
		}
		if head < tail {
			l[head], l[tail] = l[tail], l[head]
		}
		for head < tail && l[head] < l[tail] {
			head++
		}
		if head < tail {
			l[head], l[tail] = l[tail], l[head]
		}
	}
	return head
}

func merge_sort(l []int, head int, tail int) {
	tmp := make([]int, len(l))
	copy(tmp, l)
	merge_sort_i(l, head, tail, tmp)
}

func merge_sort_i(l []int, head int, tail int, tmp []int) {
	fmt.Println(l)
	switch {
	case tail-head < 1:
		return
	case tail-head == 1 && l[head] > l[tail]:
		l[head], l[tail] = l[tail], l[head]
	default:
		mid := (head + tail) / 2
		merge_sort(l, head, mid)
		merge_sort(l, mid+1, tail)
		pos1 := head
		pos2 := mid + 1
		pos_tmp := head
		for (pos1 < mid+1) || (pos2 < tail+1) {
			switch {
			case pos1 < mid+1 && pos2 < tail+1:
				if l[pos1] <= l[pos2] {
					tmp[pos_tmp] = l[pos1]
					pos1++
				} else {
					tmp[pos_tmp] = l[pos2]
					pos2++
				}
			case pos1 == mid+1:
				tmp[pos_tmp] = l[pos2]
				pos2++
			case pos2 == tail+1:
				tmp[pos_tmp] = l[pos1]
				pos1++
			}
			pos_tmp++
		}
		copy(l[head:tail+1], tmp[head:tail+1])
	}
}
