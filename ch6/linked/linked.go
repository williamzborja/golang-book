package main

import (
	"bytes"
	"fmt"
)

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	Head *Node
	Len  int
}

func (l *LinkedList) Add(value int) {
	node := &Node{}
	node.Value = value

	if l.Head == nil {
		l.Head = node
		return
	}

	iterator := l.Head
	for iterator.Next != nil {
		iterator = iterator.Next
	}

	iterator.Next = node
}

func (l *LinkedList) Remove(value int) {
	var previous *Node

	if l.Head.Value == value {
		l.Head = l.Head.Next
		return
	}

	current := l.Head
	for current != nil {
		if value == current.Value {
			previous.Next = current.Next
			return
		}
		previous = current
		current = current.Next
	}
}

func (l *LinkedList) Get(value int) *Node {
	iterator := l.Head
	for iterator != nil {
		if iterator.Value == value {
			return iterator
		}
		iterator = iterator.Next
	}
	return nil
}

func (l LinkedList) String() string {

	var buf bytes.Buffer
	current := l.Head

	for current != nil {
		fmt.Fprintf(&buf, "%d -> ", current.Value)
		current = current.Next
	}

	return buf.String()
}

func main() {
	var linked LinkedList

	linked.Add(1)
	linked.Add(2)
	linked.Remove(1)
	linked.Add(5)
	linked.Get(5)
	fmt.Println(linked)
}

func intToRoman(num int) string {
	var roman string = ""
	var numbers = []int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}
	var romans = []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"}
	var index = len(romans) - 1

	for num > 0 {
		for numbers[index] <= num {
			roman += romans[index]
			num -= numbers[index]
		}
		index -= 1
	}

	return roman
}
