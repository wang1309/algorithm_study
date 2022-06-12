package array


// https://leetcode.cn/problems/remove-duplicates-from-sorted-array/
// 删除有序数组中的重复项
// [0,0,1,1,2]
func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	slow, fast := 0, 0
	for fast < len(nums) {
		if nums[slow] != nums[fast] {
			slow++
			nums[slow] = nums[fast]
		}

		fast++
	}


	return slow+1
}


// https://leetcode.cn/problems/remove-element/
// 移除元素
func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}

	slow, fast := 0, 0
	for fast < len(nums) {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}

		fast++
	}

	return slow
}
