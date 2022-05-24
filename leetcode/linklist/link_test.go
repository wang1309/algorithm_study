package linklist

import (
	"fmt"
	"testing"
)

func TestLink(t *testing.T)  {
	head := new(ListNode)
	head.Val = 1
	ln2 := new(ListNode)
	ln2.Val = 2
	ln3 := new(ListNode)
	ln3.Val = 3
	ln4 := new(ListNode)
	ln4.Val = 4
	ln5 := new(ListNode)
	ln5.Val = 5
	head.Next = ln2
	ln2.Next = ln3
	ln3.Next = ln4
	ln4.Next = ln5

	pre := reverseList(head)
	for pre != nil {
		fmt.Println(pre.Val)
		pre = pre.Next
	}

}



