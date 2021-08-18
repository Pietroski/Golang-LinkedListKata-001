package LinkedListKata

import (
	"sync"
)

var (
	wg sync.WaitGroup
)

type SinglyLinkedListNode struct {
	data int
	next *SinglyLinkedListNode
}

func removeNodes(listHead *SinglyLinkedListNode, x int) *SinglyLinkedListNode {
	var filteredList *SinglyLinkedListNode

	gn := generateNodes(listHead, x)
	filteredList = concatNodes(gn, x)

	return filteredList
}

func generateNodes(listHead *SinglyLinkedListNode, x int) <-chan *SinglyLinkedListNode {
	out := make(chan *SinglyLinkedListNode)
	go func() {
		node := listHead
		for node != nil {
			if node.data <= x {
				out <- node
			}

			if node.next == nil {
				break
			}

			node = node.next
		}
		close(out)
	}()
	return out
}

func concatNodes(linkedList <-chan *SinglyLinkedListNode, x int) *SinglyLinkedListNode {
	var listToRelink = []*SinglyLinkedListNode{ <-linkedList }

	node := listToRelink[0]
	for i := 0; node != nil; i++ {
		nextItem := <-linkedList
		if nextItem == nil {
			break
		}
		listToRelink[i].next = nextItem
		listToRelink = append(listToRelink, nextItem)
		node = nextItem
	}

	// clean the last element's pointer
	listToRelink[len(listToRelink) - 2].next.next = nil
	return listToRelink[0]
}
