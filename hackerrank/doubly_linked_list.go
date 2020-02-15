package main

import (
	"bufio"
	"fmt"
)

type DoublyLinkedListNode struct {
	data int32
	next *DoublyLinkedListNode
	prev *DoublyLinkedListNode
}

type DoublyLinkedList struct {
	head *DoublyLinkedListNode
	tail *DoublyLinkedListNode
}

func (doublyLinkedList *DoublyLinkedList) insertNodeIntoDoublyLinkedList(nodeData int32) {
	node := &DoublyLinkedListNode{
		next: nil,
		prev: nil,
		data: nodeData,
	}

	if doublyLinkedList.head == nil {
		doublyLinkedList.head = node
	} else {
		doublyLinkedList.tail.next = node
		node.prev = doublyLinkedList.tail
	}

	doublyLinkedList.tail = node
}

func printDoublyLinkedList(node *DoublyLinkedListNode, sep string, writer *bufio.Writer) {
	for node != nil {
		fmt.Fprintf(writer, "%d", node.data)

		node = node.next

		if node != nil {
			fmt.Fprintf(writer, sep)
		}
	}
}

func sortedInsert(head *DoublyLinkedListNode, data int32) *DoublyLinkedListNode {
	if head == nil {
		llist1 := DoublyLinkedList{}
		llist1.insertNodeIntoDoublyLinkedList(data)
		return llist1.head
	}

	result := head
	current := head
	for current != nil {
		if current.data > data {
			node := &DoublyLinkedListNode{
				next: current,
				prev: current.prev,
				data: data,
			}
			if current.prev != nil {
				current.prev.next = node
			} else {
				result = node
			}
			current.prev = node
			break
		}

		if current.next == nil {
			node := &DoublyLinkedListNode{
				next: nil,
				prev: current,
				data: data,
			}
			current.next = node
			break
		}
		current = current.next
	}

	return result
}

func reversed(head *DoublyLinkedListNode) *DoublyLinkedListNode {

	current := head
	var newHead = head
	for current != nil {
		temp := current.prev
		current.prev = current.next
		current.next = temp
		current = current.prev
		if temp != nil {
			newHead = temp.prev
		}
	}

	return newHead
}

func main() {
	arr2 := []int32{1, 2, 3, 4}

	llist1 := DoublyLinkedList{}

	for i := 0; i < len(arr2); i++ {
		llist1.insertNodeIntoDoublyLinkedList(arr2[i])
	}

	//sortedInsert(llist1.head, 4)
	reversed(llist1.head)
}
