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

// 反转链表递归解法
func reverseListRec(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	last := reverseListRec(head.Next)
	head.Next.Next = head
	head.Next = nil
	return last
}

// 反转链表前 N 个节点
var successor *ListNode
func reverseN(head *ListNode, n int) *ListNode {
	if n == 1 {
		successor = head.Next
		return head
	}

	last := reverseN(head.Next, n-1)
	head.Next.Next = head
	head.Next = successor

	return last
}

// 反转链表的一部分
// 给一个索引区间 [m, n]（索引从 1 开始），仅仅反转区间中的链表元素。
func reverseBetweenRec(head *ListNode, m , n int) *ListNode {
	if m == 1 {
		// 相当于反转前 n 个元素
		return reverseN(head, n)
	}


	head.Next = reverseBetweenRec(head.Next, m-1, n-1)
	return head
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

// https://leetcode.cn/problems/merge-two-sorted-lists/
// 合并两个有序链表
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	}

	dummy := &ListNode{}

	p := dummy

	for list1 != nil && list2 != nil {
		if list1.Val > list2.Val {
			p.Next = list2
			list2 = list2.Next
		} else {
			p.Next = list1
			list1 = list1.Next
		}

		p = p.Next
	}

	if list1 != nil {
		p.Next = list1
	}

	if list2 != nil {
		p.Next = list2
	}

	return dummy.Next
}

// 返回链表的倒数第 k 个节点
func findFromEnd(head *ListNode, k int) *ListNode {
	if head == nil || k < 0 {
		return nil
	}

	p1 := head
	for i := 0; i < k; i++ {
		p1 = p1.Next
	}

	p2 := head

	for p1 != nil {
		p1 = p1.Next
		p2 = p2.Next
	}

	return p2
}


// https://leetcode.cn/problems/middle-of-the-linked-list/
// 链表的中间结点
func middleNode(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

// https://leetcode.cn/problems/palindrome-linked-list/
// 回文链表
var left *ListNode
func isPalindrome(head *ListNode) bool {
	if head == nil {
		return false
	}

	left = head
	return traverse(head)
}

func traverse(right *ListNode) bool {
	if right == nil {
		return true
	}

	res := traverse(right.Next)

	res = res && (right.Val == left.Val)
	left = left.Next
	return res
}