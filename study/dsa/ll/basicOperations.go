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
type linkedList struct {
	head *node
}

func main() {
	ll := linkedList{
		head: nil,
	}

	ll.add(4)
	ll.add(40)
	ll.traverse()
	fmt.Println(ll.String())
	ll.traverseRecur(ll.head)
}

func (ll *linkedList) add(v int) {
	n := node{
		value: v,
	}
	n.next = ll.head
	ll.head = &n
}

func (ll *linkedList) traverse() {
	var output, sep string
	for node := ll.head; node != nil; node = node.next {
		output = output + sep + strconv.Itoa(node.value)
		sep = " "
	}
	fmt.Printf("%v\n", output)
}

func (ll *linkedList) traverseRecur(n *node) {
	if n == nil {
		fmt.Println("")
		return
	}
	fmt.Printf("%d ", n.value)
	ll.traverseRecur(n.next)
}

func (ll *linkedList) String() string {
	//To string function
	//Use string builder because it is more efficient
	var sb strings.Builder
	for node := ll.head; node != nil; node = node.next {
		sb.WriteString(fmt.Sprintf("%d ", node.value))
	}
	return sb.String()
}
