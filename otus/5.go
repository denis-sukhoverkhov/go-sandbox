package main

type LinkedListNode struct {
	value int
	next  *LinkedListNode
	prev  *LinkedListNode
}

func (n LinkedListNode) Value() int {
	return n.value
}

func (n LinkedListNode) Next() *LinkedListNode {
	return n.next
}

func (n LinkedListNode) Prev() *LinkedListNode {
	return n.prev
}

type DoubleLinkedList struct {
	head *LinkedListNode
	tail *LinkedListNode
}

func (l *DoubleLinkedList) Len() int {
	currentNode := l.head
	length := 0
	for currentNode != nil {
		length++
		currentNode = currentNode.next
	}

	return length
}

func (l *DoubleLinkedList) PushFront(i int) {
	newNode := &LinkedListNode{
		value: i,
		next:  l.head,
	}

	if l.head != nil {
		l.head.prev = newNode
	} else {
		l.tail = newNode
	}

	l.head = newNode
}

func (l *DoubleLinkedList) PushBack(i int) {
	newNode := &LinkedListNode{
		value: i,
		prev:  l.tail,
	}

	if l.tail != nil {
		l.tail.next = newNode
	} else {
		l.head = newNode
	}

	l.tail = newNode
}

func (l *DoubleLinkedList) Last() interface{} {
	return l.tail
}

func (l *DoubleLinkedList) First() interface{} {
	return l.head
}

func (l *DoubleLinkedList) Remove(i int) {
	currentNode := l.head

	for currentNode != nil {
		if currentNode.Value() == i {
			currentNode.prev.next = currentNode.Next()
			currentNode.next.prev = currentNode.Prev()
			break
		}
		currentNode = currentNode.next
	}

}
