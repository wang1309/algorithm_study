package linklist

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * https://leetcode.cn/problems/reverse-linked-list/
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}

	return prev
}

// https://leetcode.cn/problems/linked-list-cycle/
// 环形链表
// hasCycleHash hash 表法
func hasCycleHash(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	visit := make(map[*ListNode]struct{})

	for head != nil {
		if _, ok := visit[head]; ok {
			return true
		} else {
			visit[head] = struct{}{}
		}
		head = head.Next
	}

	return false
}

// hasCycle 龟兔赛跑解法
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow, fast := head, head.Next

	for slow != fast {
		if fast == nil || fast.Next == nil {
			return false
		}

		slow = slow.Next
		fast = fast.Next.Next
	}

	return true
}

// https://leetcode.cn/problems/reverse-linked-list-ii/
// 反转链表 II
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	// 设置 dummyNode 是这一类问题的一般做法
	dummyNode := &ListNode{Val: -1}
	dummyNode.Next = head
	pre := dummyNode
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}

	cur := pre.Next
	for i := 0; i < right-left; i++ {
		next := cur.Next
		cur.Next = next.Next
		next.Next = pre.Next
		pre.Next = next
	}
	return dummyNode.Next
}

// https://leetcode.cn/problems/palindrome-linked-list/
// 回文链表
// [1,2,2,1]

// 方法一：将值复制到数组中后用双指针法
func isPalindrome1(head *ListNode) bool {
	if head == nil {
		return false
	}

	arr := []int{}

	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}

	length := len(arr)
	j := length - 1

	for i := 0; i < length/2; i++ {
		if arr[i] != arr[j] {
			return false
		}
		j--
	}

	return true
}

// https://leetcode.cn/problems/remove-nth-node-from-end-of-list/
// 删除链表的倒数第 N 个结点
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, head}
	fast, slow := head, dummy
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	for ; fast != nil; fast = fast.Next {
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return dummy.Next
}