package leetcode

import (
	"testing"
)

func Test_binary_find_1(t *testing.T) {
	nums := []int{2,3,4,5,6,7,8,10,23,56,234}
	target := 13333
	answer := binary_find(nums, target, 0, len(nums)-1)
	if(answer != 10) {
		t.Fatal("error!, answer is 7, actual is ", answer)
	}
}

func Test_lis_1(t *testing.T) {
	nums := []int{2,2}
	actual := lengthOfLIS(nums)
	if(actual != 1) {
		t.Fatal("error!, answer is 7, actual is ", actual)
	}
}
