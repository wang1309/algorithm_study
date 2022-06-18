package other

import "sort"

type Node struct {
	val   int
	index int
}

// https://leetcode.cn/problems/advantage-shuffle/
// 优势洗牌
func advantageCount(nums1 []int, nums2 []int) []int {
	var (
		res     []int  = make([]int, len(nums1))
		left    int    = 0
		right   int    = len(nums1) - 1
		nodeArr []Node = make([]Node, len(nums1))
	)

	sort.Slice(nums1, func(i, j int) bool {
		return nums1[i]-nums1[j] < 0
	})

	for i, n := range nums2 {
		nodeArr[i] = Node{val: n, index: i}
	}

	sort.Slice(nodeArr, func(i, j int) bool {
		return nodeArr[i].val-nodeArr[j].val < 0
	})

	for i := len(nodeArr) - 1; i >= 0; i-- {
		val := nodeArr[i].val
		index := nodeArr[i].index
		if nums1[right] > val {
			res[index] = nums1[right]
			right--
		} else {
			res[index] = nums1[left]
			left++
		}
	}

	return res
}
