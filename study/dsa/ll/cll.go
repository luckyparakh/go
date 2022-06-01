package main

import (
	"fmt"
	"strconv"
	"strings"
)

type node struct {
	value int
	next  *node
}
type circularLL struct {
	head *node
}

func (cll *circularLL) CreateNode(val int) *node {
	return &node{
		value: val,
	}
}

func (cll *circularLL) traversal() {
	if cll.head == nil {
		return
	}
	sb := strings.Builder{}
	sb.WriteString(strconv.Itoa(cll.head.value)+" ")
	tmpNode := cll.head.next

	for tmpNode != cll.head {
		sb.WriteString(strconv.Itoa(tmpNode.value)+" ")
		tmpNode = tmpNode.next
	}
	fmt.Println(sb.String())
}
func main() {
	cll := circularLL{}
	n1 := cll.CreateNode(5)
	n2 := cll.CreateNode(50)
	n3 := cll.CreateNode(500)
	cll.head = n1
	n1.next = n2
	n2.next = n3
	n3.next = cll.head
	cll.traversal()
}
