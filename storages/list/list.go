package list

import (
	"fmt"
	"strings"
)

type List struct {
	indexes   int64
	len       int64
	firstNode *Node
	LastNode  *Node
}

func NewList() *List {
	return &List{
		indexes:   0,
		len:       0,
		firstNode: nil,
		LastNode:  nil,
	}
}

func (l *List) Add(data int64) (index int64) {
	if l.firstNode == nil {
		l.indexes += 1
		n := &Node{
			index:    l.indexes,
			data:     data,
			nextNode: nil,
		}
		l.firstNode = n
		l.LastNode = n
		l.len++
		return n.index
	}
	l.indexes += 1

	newNode := &Node{
		index:    l.indexes,
		data:     data,
		nextNode: nil,
	}

	l.LastNode.nextNode = newNode
	l.LastNode = newNode
	l.len++
	return newNode.index
}

func (l *List) Get(index int64) (data int64) {
	for n := l.firstNode; n != nil; n = n.nextNode {
		if n.index == index {
			return n.data
		}
	}
	return
}

func (l *List) Len() int64 {
	return l.len
}

func (l *List) Delete(index int64) {

	if index == l.firstNode.index {
		l.firstNode = l.firstNode.nextNode
		l.len--
		return
	}

	temp1 := l.firstNode.nextNode
	temp2 := l.firstNode

	for ; temp1 != nil; temp1 = temp1.nextNode {
		if index == temp1.index {
			if temp1 == l.LastNode {
				l.LastNode = temp2
			}
			temp2.nextNode = temp1.nextNode
			return
		}
		temp2 = temp2.nextNode
	}
}

func (l *List) String() string {
	if l.firstNode == nil {
		return "empty list"
	}
	sb := strings.Builder{}
	n := l.firstNode
	for {
		sb.WriteString(fmt.Sprintf("index: '%d' \t data: '%d'", n.index, n.data))
		if n.nextNode == nil {
			return sb.String()
		}
		sb.WriteString("\n")
		n = n.nextNode
	}
}

func (l *List) SortIncrease() {
	for i := int64(1); i < l.len; i++ {
		first := l.firstNode
		second := l.firstNode.nextNode
		for j := int64(1); j < l.len; j++ {
			if first.data > second.data {
				temp1 := first.data
				temp2 := first.index
				first.index = second.index
				first.data = second.data
				second.data = temp1
				second.index = temp2

			}
			first = first.nextNode
			second = second.nextNode
		}
	}

}

func (l *List) SortDecrease() {

}
