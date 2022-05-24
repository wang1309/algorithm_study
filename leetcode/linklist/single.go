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
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
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


// hasCycle 龟兔赛跑法
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