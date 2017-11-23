package leetcode

import "testing"

func Test_quicksort_1(t *testing.T) {
	l := []int{5, 3, 6, 2, 15, 30, 7, 4}
	quick_sort(l, 0, len(l)-1)
	result := []int{2, 3, 4, 5, 6, 7, 15, 30}
	if len(l) != len(result) {
		t.Fatal("length not equal!")
	}
	for index, _ := range l {
		if l[index] != result[index] {
			t.Fatalf("%d pos val not equal! It should be %d, but %d", index, result[index], l[index])
		}
	}
}

func Test_mergesort_1(t *testing.T) {
	l := []int{5, 3, 6, 2, 15, 30, 7, 4}
	merge_sort(l, 0, len(l)-1)
	result := []int{2, 3, 4, 5, 6, 7, 15, 30}
	if len(l) != len(result) {
		t.Fatal("length not equal!")
	}
	for index, _ := range l {
		if l[index] != result[index] {
			t.Fatalf("%d pos val not equal! It should be %d, but %d", index, result[index], l[index])
		}
	}
}
