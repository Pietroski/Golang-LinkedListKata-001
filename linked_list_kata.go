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
		for node.next != nil {
			if node.data <= x {
				out <- node
			}

			node = node.next
		}
		close(out)
	}()
	return out
}

func concatNodes(linkedList <-chan *SinglyLinkedListNode, x int) *SinglyLinkedListNode {
	var listToRelink = []*SinglyLinkedListNode{ <-linkedList }

	node := listToRelink[0].next
	for node != nil {
		if node.data <= x {
			listToRelink = append(listToRelink, node)
		}
		node = node.next
	}

	listToRelink[len(listToRelink) - 2].next.next = nil
	return listToRelink[0]
}
