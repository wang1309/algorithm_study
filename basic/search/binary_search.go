package search

func BinarySearch(nums []int, val int) int {
	if len(nums) == 0 {
		return -1
	}

	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)>>2

		if nums[mid] == val {
			return val
		} else if nums[mid] > val {
			left = mid + 1
		} else if nums[mid] < val {
			right = mid - 1
		}
	}

	return -1
}
