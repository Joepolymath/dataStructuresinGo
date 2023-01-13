package main

import "fmt"

// IMPLEMENTING LINKED LIST IN GOLANG...
type Node struct {
	value int
	next *Node
}

type LinkedList struct {
	head *Node
	len int
}

func (l *LinkedList) isEmpty() bool {
	return l.len == 0
}

func (l *LinkedList) insertToEnd(val int) Node {
	n := Node{}
	n.value = val

	if l.len == 0 {
		l.head = &n
		l.len++
		return n
	}
	currentHead := l.head
	for i:= 0; i < l.len; i++ {
		if currentHead.next == nil {
			currentHead.next = &n
			l.len++
			return n
		}
		currentHead = currentHead.next
	}
	return n
} 

func (l *LinkedList) insertToFront(val int) {
	n := Node{}
	n.value = val
	if l.len == 0 {
		l.head = &n
		l.len++
		return
	}
	temp := l.head
	l.head = &n
	l.head.next = temp
	l.len++
	return
}

func main()  {
	list := LinkedList{}
	fmt.Println(list.isEmpty())
	fmt.Println(list)
	list.insertToEnd(5)
	list.insertToFront(2)
	list.insertToEnd(6)
	fmt.Println(list)
	fmt.Println(list.isEmpty())
}