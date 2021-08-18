package LinkedListKata

import (
	"fmt"
	"testing"
)

type testModel struct {
	x int
	input *SinglyLinkedListNode
	output *SinglyLinkedListNode
}

var tableTestLLMap = map[int]testModel{
	0: {
		x: 3,
		input:  &SinglyLinkedListNode{
			data: 5,
			next: &SinglyLinkedListNode{
				data: 1,
				next: &SinglyLinkedListNode{
					data: 2,
					next: &SinglyLinkedListNode{
						data: 3,
						next: &SinglyLinkedListNode{
							data: 4,
							next: &SinglyLinkedListNode{
								data: 5,
								next: nil,
							},
						},
					},
				},
			},
		},
		output: &SinglyLinkedListNode{
			data: 1,
			next: &SinglyLinkedListNode{
				data: 2,
				next: &SinglyLinkedListNode{
					data: 3,
					next: nil,
				},
			},
		},
	},
}

func TestRemoveNodes(t *testing.T) {
	fmt.Println("TestRemoveNodes")

	for testNum, test := range tableTestLLMap {
		x := test.x
		input := test.input
		output := test.output

		fro := removeNodes(input, x)

		node := fro
		out := output
		for node != nil {

			// shallow checking!!
			if node.data != out.data {
				msg := fmt.Errorf("test number %v failed | expected: %v, got: %v", testNum, out, node)
				panic(msg)
			}

			node = node.next
			out = out.next
		}
	}

	fmt.Println("")
}
