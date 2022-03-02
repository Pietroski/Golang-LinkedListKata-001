package LinkedListKata

type SinglyLinkedListNode struct {
	data int
	next *SinglyLinkedListNode
}

func removeNodes(llist *SinglyLinkedListNode, position int) *SinglyLinkedListNode {
	if position == 0 {
		return llist.next
	}

	filterNode(llist, position, 0)
	return llist
}

func filterNode(llist *SinglyLinkedListNode, position, counter int) *SinglyLinkedListNode {
	if llist != nil && llist.next != nil {
		if counter+1 == position {
			llist.next = llist.next.next
		}

		counter++
		return filterNode(llist.next, position, counter)
	}

	return llist
}
