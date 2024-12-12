package structs

import (
	"fmt"
)

type element struct {
	Data interface{}
	Next *element
}

type linkedList struct {
	Head *element
	len  int
}

func NewLinkedList() linkedList {
	return linkedList{}
}

func (list *linkedList) AddToHead(data interface{}) {
	newElement := &element{
		Data: data,
		Next: nil,
	}

	if list.Head == nil {
		list.Head = newElement
		list.len++
		return
	}
	newElement.Next = list.Head
	list.Head = newElement
	list.len++
}

func (list *linkedList) AddToTail(data interface{}) {
	newElement := &element{
		Data: data,
		Next: nil,
	}

	if list.Head == nil {
		list.Head = newElement
		list.len++
		return
	}
	current := list.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newElement
	list.len++
}

func (list *linkedList) RemoveHead() error {

	if list.Head == nil {
		return fmt.Errorf("No elements in Linked List")
	}
	list.Head = list.Head.Next
	list.len--
	return nil
}

func (list *linkedList) RemoveTail() error {

	if list.Head == nil {
		return fmt.Errorf("No elements in Linked List")
	}
	if list.Head.Next == nil {
		list.Head = nil
		list.len--
		return nil
	}
	current := list.Head
	for current.Next.Next != nil {
		current = current.Next
	}
	current.Next = nil
	list.len--
	return nil
}

func (list *linkedList) Remove(el any) error {

	if list.Head == nil {
		return fmt.Errorf("No elements in Linked List")
	}
	current := list.Head
	var prev *element
	for {
		if current.Data == el {
			if prev == nil {
				current.Data = nil
				list.len--
				return nil
			}
			current.Data = nil
			prev.Next = current.Next
			list.len--
			return nil
		}
		if current.Next == nil {
			return fmt.Errorf("No such element in List")
		}
		prev = current
		current = current.Next
	}

}

func (list *linkedList) Traverse() error {
	if list.Head == nil {
		return fmt.Errorf("TranverseError: List is empty")
	}
	current := list.Head
	for current != nil {
		fmt.Println(current.Data)
		current = current.Next
	}
	return nil
}
