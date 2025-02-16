package dsa_test

import (
	"log/slog"
	"testing"

	"github.com/stretchr/testify/assert"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head

	for curr != nil {
		next := curr.Next // Store next node
		curr.Next = prev  // Reverse current node's pointer
		prev = curr       // Move prev to current node
		curr = next       // Move to next node
	}
	return prev
}

// Helper function to print the linked list
func PrintList(head *ListNode) {
	for head != nil {
		slog.Info("Node:", "val", head.Val)
		head = head.Next
	}
}

func TestReverse(t *testing.T) {
	t.Parallel()
	// Creating a linked list: 1 -> 2 -> 3 -> 4 -> nil
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}

	slog.Info("Original list:", "list", head)
	assert.Equal(t, 1, head.Val)

	PrintList(head)

	reversedHead := ReverseList(head)

	slog.Info("Reversed list:", "list", reversedHead)
	PrintList(reversedHead)
	assert.Equal(t, 4, reversedHead.Val)
}

type LinkedList struct {
	Head *ListNode
}

func (l *LinkedList) Read(index int) int {
	current := l.Head
	for i := 0; i < index; i++ {
		slog.Info("Current node:", "val", current.Val)
		current = current.Next
	}
	return current.Val
}

func TestRead(t *testing.T) {
	t.Parallel()
	// Creating a linked list: 1 -> 2 -> 3 -> 4 -> nil
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}
	ll := LinkedList{head}

	slog.Info("Reading index 2:", "val", ll.Read(2))
	assert.Equal(t, 3, ll.Read(2))
}

func (l *LinkedList) IndexOf(value int) int {
	current := l.Head
	index := 0
	for current != nil {
		if current.Val == value {
			return index
		}
		current = current.Next
		index++
	}
	return -1
}

func TestIndexOf(t *testing.T) {
	t.Parallel()
	// Creating a linked list: 1 -> 2 -> 3 -> 4 -> nil
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}
	ll := LinkedList{head}

	slog.Info("Index of 3:", "index", ll.IndexOf(3))
	assert.Equal(t, 2, ll.IndexOf(3))
}

func (l *LinkedList) Insert(index, value int) {
	newNode := &ListNode{Val: value}
	if index == 0 {
		newNode.Next = l.Head
		l.Head = newNode
		return
	}

	current := l.Head
	for i := 0; i < index-1; i++ {
		current = current.Next
	}
	newNode.Next = current.Next
	current.Next = newNode
}

func TestInsert(t *testing.T) {
	t.Parallel()
	// Creating a linked list: 1 -> 2 -> 3 -> 4 -> nil
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}
	ll := LinkedList{head}

	ll.Insert(2, 5)
	assert.Equal(t, 5, ll.Read(2))
}

func (l *LinkedList) Delete(index int) {
	if index == 0 {
		l.Head = l.Head.Next
		return
	}

	current := l.Head
	for i := 0; i < index-1; i++ {
		current = current.Next
	}
	current.Next = current.Next.Next
}

func TestDelete(t *testing.T) {
	t.Parallel()
	// Creating a linked list: 1 -> 2 -> 3 -> 4 -> nil
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}
	ll := LinkedList{head}

	ll.Delete(2)
	assert.Equal(t, 4, ll.Read(2))
}
