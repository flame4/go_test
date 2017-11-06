package leetcode

import "fmt"

func nextPermutation(nums []int)  {
	// 找到最后一个递减子序列
	last_index := 0
	for index, _ := range nums {
		if index < len(nums) - 1 {
			if nums[index] < nums[index+1] {
				last_index = index + 1
			}
		}
	}

	// 将最后一个递减子序列前面的数字和这个序列中第一个比它大的换一下
	candicate_index := last_index - 1
	target_index := last_index
	if(candicate_index >= 0) {
		for target_index < len(nums) {
			if(nums[candicate_index] < nums[target_index]) {
				target_index++
			} else {
				break
			}
		}
		target_index--
		nums[target_index], nums[candicate_index] = nums[candicate_index], nums[target_index]
	}

	// 将最后一个子序列正序排列
	head := last_index
	tail := len(nums) - 1
	for head < tail {
		nums[head], nums[tail] = nums[tail], nums[head]
		head++
		tail--
	}

}