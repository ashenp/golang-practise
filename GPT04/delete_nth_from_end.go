package GPT02

type ListNode struct {
	Val  int
	Next *ListNode
}

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Val: 0, Next: head}
	slow, fast := dummy, head
	for i := 0; i < n && fast != nil; i++ {
		fast = fast.Next
	}

	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}

	if slow.Next != nil {
		slow.Next = slow.Next.Next
	}
	return dummy.Next
}
