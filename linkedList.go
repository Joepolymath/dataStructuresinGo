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

func (l *LinkedList) insert(val int) Node {
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

func main()  {
	list := LinkedList{}
	fmt.Println(list.isEmpty())
	fmt.Println(list)
	list.insert(5)
	list.insert(6)
	fmt.Println(list)
	fmt.Println(list.isEmpty())
}