package LinkedListKata

type SinglyLinkedListNode struct {
	data int
	next *SinglyLinkedListNode
}

func removeNodes(listHead *SinglyLinkedListNode, x int) *SinglyLinkedListNode {
	var filteredList []*SinglyLinkedListNode

	node := listHead
	for node.next != nil {
		if node.data <= x {
			filteredList = append(filteredList, node)
		}

		node = node.next
	}

	return concatNodes(filteredList)
}

func concatNodes(linkedList []*SinglyLinkedListNode) *SinglyLinkedListNode {
	var listToRelink []*SinglyLinkedListNode

	linkedLen := len(linkedList)
	for i := 1; i < linkedLen; i++ {
		ll := linkedList[i-1]
		ll.next = linkedList[i]
		listToRelink = append(listToRelink, ll)
	}

	listToRelink[linkedLen - 2].next.next = nil
	return listToRelink[0]
}
