package leetcode

func lengthOfLIS(nums []int) int {
	if(len(nums) == 0) {
		return 0
	}
	endings := make([]int, len(nums))
	length := 1
	endings[0] = nums[0]
	for _, v := range nums[1:] {
		pos := binary_find(endings, v, 0, length-1)
		endings[pos+1] = v
		if(pos == length-1) {
			length++
		}
	}
	return length
}


// 返回在nums中 < target的那个index, -1 ~ len(nums)-1
func binary_find(nums []int, target int, head int, tail int) int {

	for head <= tail {
		mid := (head + tail) / 2
		if(nums[mid] >= target) {
			tail = mid - 1
		} else if nums[mid] < target {
			head = mid + 1
		}
	}
	if head > tail {
		return tail
	} else {
		return head
	}
}