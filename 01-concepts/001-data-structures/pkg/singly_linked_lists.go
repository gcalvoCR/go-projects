package pkg

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (list *LinkedList) Insert(data int) {
	newNode := &Node{data: data}

	if list.head == nil {
		list.head = newNode
		return
	}
	current := list.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}

// Append an item to the end of the list
func (l *LinkedList) Append(number int) {
	// Create the new node
	node := &Node{data: number, next: nil}

	// 2 cases
	if l.head == nil {
		// When list empty
		l.head = node
	} else {
		// When list has values
		p := l.head
		for p.next != nil {
			p = p.next
		}
		p.next = node
	}
}

func (l *LinkedList) Prepend(number int) {
	// Create the node
	node := &Node{data: number, next: nil}

	// 2 cases
	// When list is empty
	if l.head == nil {
		l.head = node
	} else {
		temp := l.head
		node.next = temp
		l.head = node
	}
}

func (l LinkedList) Length() int {
	p := l.head
	var len int
	for p != nil {
		len++
		p = p.next
	}
	return len
}

func (l LinkedList) Contains(number int) bool {
	p := l.head
	for p != nil {
		if p.data == number {
			return true
		}
		p = p.next
	}
	return false
}

func (l LinkedList) Index(number int) int {
	p := l.head
	var index int
	var isFound bool
	for p != nil {
		if p.data == number {
			isFound = true
			break
		}
		index++
		p = p.next
	}
	if isFound {
		return index
	}
	return -1
}

func (l *LinkedList) DeleteFirst() {
	if l.head != nil {
		l.head = l.head.next
	}
}

func (l *LinkedList) Delete(data int) {
	if l.head != nil {
		if l.head.data == data {
			l.head = l.head.next
		} else {
			curr := l.head
			prev := &Node{}
			var isFound bool
			for curr.next != nil {
				if curr.data == data {
					isFound = true
					break
				}
				prev = curr
				curr = curr.next
			}
			if isFound {
				prev.next = curr.next
			}
		}
	}
}

func (list *LinkedList) Display() {
	current := list.head

	if current == nil {
		fmt.Println("Linked list is empty")
		return
	}

	fmt.Print("Linked list: ")
	for current != nil {
		fmt.Printf("%d ", current.data)
		current = current.next
	}
	fmt.Println()
}
