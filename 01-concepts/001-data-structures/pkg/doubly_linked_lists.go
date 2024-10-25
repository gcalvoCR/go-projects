package pkg

import (
	"fmt"

	"github.com/fatih/color"
)

type DoublyNode struct {
	data int
	prev *DoublyNode
	next *DoublyNode
}

type DoublyLinkedList struct {
	head *DoublyNode
	tail *DoublyNode
}

// Insert at head
func (list *DoublyLinkedList) InsertAtBeginning(data int) {
	newNode := &DoublyNode{data: data}

	if list.head == nil {
		list.head = newNode
		list.tail = newNode
		return
	}

	// link the new node
	newNode.next = list.head
	list.head.prev = newNode

	// assign the new node
	list.head = newNode
}

// Insert at tail
func (list *DoublyLinkedList) InsertAtEnd(data int) {
	newNode := &DoublyNode{data: data}

	if list.head == nil {
		list.head = newNode
		list.tail = newNode
		return
	}

	newNode.prev = list.tail
	list.tail.next = newNode
	list.tail = newNode
}

func (list *DoublyLinkedList) TraverseForward() {
	current := list.head

	if current == nil {
		fmt.Println("Linked list is empty")
		return
	}

	fmt.Print("Linked list: ")
	for current != nil {
		prev := -1
		if current.prev != nil {
			prev = current.prev.data
		}
		fmt.Printf(" <-[ %d ", prev)
		fmt.Printf("%s", color.GreenString(fmt.Sprintf("| %d |", current.data)))
		next := -1
		if current.next != nil {
			next = current.next.data
		}
		fmt.Printf(" %d ]-> ", next)

		current = current.next
	}
	fmt.Println()
}

func (list *DoublyLinkedList) TraverseBackwards() {
	current := list.tail

	if current == nil {
		fmt.Println("Linked list is empty")
		return
	}

	fmt.Print("Linked list: ")
	for current != nil {
		prev := -1
		if current.prev != nil {
			prev = current.prev.data
		}
		fmt.Printf(" <-[ %d ", prev)
		fmt.Printf("%s", color.GreenString(fmt.Sprintf("| %d |", current.data)))
		next := -1
		if current.next != nil {
			next = current.next.data
		}
		fmt.Printf(" %d ]-> ", next)

		current = current.next
	}
	fmt.Println()
}
