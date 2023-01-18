package main

import "fmt"

type Node struct {
	Val  interface{}
	Next *Node
}

type LinkedList struct {
	Length int
	Head   *Node
	Tail   *Node
}

func (l LinkedList) Len() int {
	return l.Length
}

func (l LinkedList) Display() {
	for l.Head != nil {
		fmt.Printf("%v ->", l.Head.Val)
		l.Head = l.Head.Next
	}
	fmt.Println()
}

func (l *LinkedList) PushBack(n *Node) {
	if l.Head == nil {
		l.Head = n
		l.Tail = n
		l.Length++
	} else {
		l.Tail.Next = n
		l.Tail = n
		l.Length++
	}
}

func (l *LinkedList) RemoveFromStart() {
	if l.Head == nil {
		fmt.Println("Linked list is empty")
		return
	}
	l.Head = l.Head.Next
	return
}

func (l *LinkedList) RemoveFromEnd() {
	if l.Head == nil {
		fmt.Println("Linked list is empty")
		return
	}
	temp := l.Head
	for temp.Next.Next != nil {
		temp = temp.Next
	}
	temp.Next = nil
	return
}

func (l *LinkedList) RemoveFromBetween(data interface{}) {
	if l.Head == nil {
		fmt.Println("Linked list is empty")
		return
	}
	temp := l.Head
	for temp.Next != nil && temp.Next.Val != data {
		temp = temp.Next
	}
	temp.Next = temp.Next.Next
	return
}

func CreateNode(data interface{}) *Node {
	return &Node{Val: data}
}

func main() {
	list := &LinkedList{}
	list.RemoveFromStart()
	list.PushBack(CreateNode(10))
	list.PushBack(CreateNode(20))
	list.PushBack(CreateNode("str"))
	list.PushBack(CreateNode(10.01))
	list.PushBack(CreateNode("asd"))
	list.PushBack(CreateNode(45))
	list.Display()
	list.RemoveFromStart()
	list.RemoveFromEnd()
	list.RemoveFromBetween("str")
	list.RemoveFromBetween(10.01)
	list.RemoveFromBetween(34)
	list.Display()
}
