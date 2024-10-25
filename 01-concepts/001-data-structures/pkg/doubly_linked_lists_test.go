package pkg_test

import (
	"testing"

	"github.com/gcalvoCR/data-structures-go/pkg"
)

func Test_DoublyLinkedList(t *testing.T) {

	dll := pkg.DoublyLinkedList{}
	dll.InsertAtBeginning(10)
	dll.InsertAtBeginning(20)
	dll.InsertAtBeginning(30)
	dll.InsertAtBeginning(40)
	dll.TraverseForward()

	dll2 := pkg.DoublyLinkedList{}
	dll2.InsertAtEnd(10)
	dll2.InsertAtEnd(20)
	dll2.InsertAtEnd(30)
	dll2.InsertAtEnd(40)
	dll2.TraverseForward()

}
