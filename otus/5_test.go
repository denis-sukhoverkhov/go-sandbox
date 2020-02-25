package main

import (
	"reflect"
	"testing"
)

func TestLinkedList(t *testing.T) {
	_ = ` 
		List // тип контейнер 
		Len() // длинна списка 
		First() // первый Item 
		Last() // последний Item 
		PushFront(v interface{}) // добавить значение в начало 
		PushBack(v interface{}) // добавить значение в конец 
		Remove(i Item) // удалить элемент?
		Item // элемент списка 
		Value() interface{} // возвращает значение 
		Nex() *Item // следующий Item 
		Prev() *Item // предыдущий `

	dll := DoubleLinkedList{}
	length := dll.Len()
	expected := 0
	if !reflect.DeepEqual(length, expected) {
		t.Errorf("Len is not correct, got: %v, expected: %v", length, expected)
	}

	dll.PushFront(14)
	dll.PushFront(17)
	length = dll.Len()
	expected = 2
	if !reflect.DeepEqual(dll.Len(), expected) {
		t.Errorf("Len is not correct, got: %v, expected: %v", length, expected)
	}

	dll.PushBack(11)

	expectedLastNode := &LinkedListNode{
		value: 11,
		prev:  dll.tail.prev,
	}
	if !reflect.DeepEqual(dll.Last(), expectedLastNode) {
		t.Errorf("Last is not correct, got: %v, expected: %v", dll.Last(), expectedLastNode)
	}

	expectedFirstNode := &LinkedListNode{
		value: 17,
		next:  dll.head.next,
	}
	if !reflect.DeepEqual(dll.First(), expectedFirstNode) {
		t.Errorf("First is not correct, got: %v, expected: %v", dll.First(), expectedFirstNode)
	}

	dll.PushBack(11)
	dll.PushBack(55)
	dll.PushBack(34)
	dll.PushBack(67)

	dll.Remove(11)

	anotherDll := DoubleLinkedList{}
	anotherDll.PushFront(14)
	anotherDll.PushFront(17)
	anotherDll.PushBack(11)
	anotherDll.PushBack(55)
	anotherDll.PushBack(34)
	anotherDll.PushBack(67)
	if !reflect.DeepEqual(dll, anotherDll) {
		t.Errorf("Not equal linked lists, got: %v, expected: %v", dll, anotherDll)
	}

}
