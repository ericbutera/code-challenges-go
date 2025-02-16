package dsa_test

// Use container/list.

type DoublyListNode struct {
	Val  int
	Prev *DoublyListNode
	Next *DoublyListNode
}

// DoublyLinkedList is a doubly linked list
type DoublyLinkedList struct {
	Head *DoublyListNode
	Tail *DoublyListNode
}

// NewDoublyLinkedList creates a new doubly linked list
func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{}
}

// AddAtHead adds a node of value val before the first element of the doubly linked list
