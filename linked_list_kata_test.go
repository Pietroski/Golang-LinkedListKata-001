package LinkedListKata

import (
	"fmt"
	"testing"
)

type testModel struct {
	position int
	input    *SinglyLinkedListNode
	output   *SinglyLinkedListNode
}

var tableTestLLMap = map[int]testModel{
	0: {
		position: 2,
		input: &SinglyLinkedListNode{
			data: 1,
			next: &SinglyLinkedListNode{
				data: 2,
				next: &SinglyLinkedListNode{
					data: 3,
					next: &SinglyLinkedListNode{
						data: 4,
						next: &SinglyLinkedListNode{
							data: 5,
							next: &SinglyLinkedListNode{
								data: 6,
								next: &SinglyLinkedListNode{
									data: 7,
									next: nil,
								},
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
					data: 4,
					next: &SinglyLinkedListNode{
						data: 5,
						next: &SinglyLinkedListNode{
							data: 6,
							next: &SinglyLinkedListNode{
								data: 7,
								next: nil,
							},
						},
					},
				},
			},
		},
	},
	1: {
		position: 6,
		input: &SinglyLinkedListNode{
			data: 1,
			next: &SinglyLinkedListNode{
				data: 2,
				next: &SinglyLinkedListNode{
					data: 3,
					next: &SinglyLinkedListNode{
						data: 4,
						next: &SinglyLinkedListNode{
							data: 5,
							next: &SinglyLinkedListNode{
								data: 6,
								next: &SinglyLinkedListNode{
									data: 7,
									next: nil,
								},
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
					next: &SinglyLinkedListNode{
						data: 4,
						next: &SinglyLinkedListNode{
							data: 5,
							next: &SinglyLinkedListNode{
								data: 6,
								next: nil,
							},
						},
					},
				},
			},
		},
	},
	2: {
		position: 0,
		input: &SinglyLinkedListNode{
			data: 1,
			next: &SinglyLinkedListNode{
				data: 2,
				next: &SinglyLinkedListNode{
					data: 3,
					next: &SinglyLinkedListNode{
						data: 4,
						next: &SinglyLinkedListNode{
							data: 5,
							next: &SinglyLinkedListNode{
								data: 6,
								next: &SinglyLinkedListNode{
									data: 7,
									next: nil,
								},
							},
						},
					},
				},
			},
		},
		output: &SinglyLinkedListNode{
			data: 2,
			next: &SinglyLinkedListNode{
				data: 3,
				next: &SinglyLinkedListNode{
					data: 4,
					next: &SinglyLinkedListNode{
						data: 5,
						next: &SinglyLinkedListNode{
							data: 6,
							next: &SinglyLinkedListNode{
								data: 7,
								next: nil,
							},
						},
					},
				},
			},
		},
	},
}

func TestRemoveNodes(t *testing.T) {
	fmt.Println("TestRemoveNodes")

	for testNum, test := range tableTestLLMap {
		position := test.position
		input := test.input
		output := test.output

		fro := removeNodes(input, position)

		node := fro
		out := output
		for node != nil {
			// shallow checking!!
			if node.data != out.data {
				msg := fmt.Errorf("test number %v failed | expected: %v, got: %v", testNum, out, node)
				t.Error(msg)
			}

			node = node.next
			out = out.next
		}
	}

	fmt.Println("")
}
