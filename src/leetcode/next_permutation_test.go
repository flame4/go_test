package leetcode

import (
	"testing"
)

func Test_next_permutation_1(t *testing.T)  {
	// nums := []int{1,2,3,4,5}
	// nums_t := []int{1,2,3,5,4}
	// nextPermutation(nums)
	// compareList(nums_t, nums, t)
	nums := []int{3,1,4,6,2}
	nums_t := []int{3,1,6,2,4}
	nextPermutation(nums)
	compareList(nums_t, nums, t)

}

func compareList(nums []int, nums2 []int, t *testing.T) {
	if(len(nums) != len(nums2)) {
		t.Fatal("length not Equal!")
	}
	for index, _ := range(nums) {
		if(nums[index] != nums2[index]) {
			t.Fatal("not equal!", "answer: ", nums, "actual:", nums2)
		}
	}
}
