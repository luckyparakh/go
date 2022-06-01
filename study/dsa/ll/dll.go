package main

import (
	"fmt"
	"strconv"
	"strings"
)

type node struct {
	value int
	prev  *node
	next  *node
}

type doubleLL struct {
	head *node
}

func (dll *doubleLL) CreateNode(val int) *node {
	return &node{
		value: val,
	}
}
func (dll *doubleLL) InsertBeing(val int) {
	newNode := dll.CreateNode(val)
	newNode.next = dll.head
	if dll.head != nil {
		dll.head.prev = newNode
	}
	dll.head = newNode
}

func (dll *doubleLL) InsertEnd(val int) {
	newNode := dll.CreateNode(val)
	if dll.head == nil {
		dll.head = newNode
		return
	}
	tempNode := dll.head
	for tempNode.next != nil {
		tempNode = tempNode.next
	}
	tempNode.next = newNode
	newNode.prev = tempNode
}
func (dll *doubleLL) String() string {
	sb := strings.Builder{}
	for tempNode := dll.head; tempNode != nil; tempNode = tempNode.next {
		sb.WriteString(strconv.Itoa(tempNode.value) + " ")
	}
	return sb.String()
}
func (dll *doubleLL) reverse() {
	if dll.head == nil || dll.head.next == nil {
		return
	}
	tmpNode := dll.head
	for tmpNode.next != nil {
		tmpNode.prev, tmpNode.next = tmpNode.next, tmpNode.prev
		tmpNode = tmpNode.prev
	}
	tmpNode.prev, tmpNode.next = tmpNode.next, tmpNode.prev
	dll.head = tmpNode
}
func (dll *doubleLL) delHead() {
	if dll.head == nil {
		return
	}
	if dll.head.next == nil {
		dll.head = nil
		return
	}
	nextNode := dll.head.next
	nextNode.prev = nil
	dll.head.next = nil
	dll.head = nextNode

	//simple
	// dll.head=dll.head.next
	// dll.head.prev=nil

}
func (dll *doubleLL) delEnd() {
	if dll.head == nil {
		return
	}
	if dll.head.next == nil {
		dll.head = nil
		return
	}
	tmpNode := dll.head
	for tmpNode.next != nil {
		tmpNode = tmpNode.next
	}
	prevNode := tmpNode.prev
	prevNode.next = nil
	tmpNode.prev = nil

}
func main() {
	dll := doubleLL{}
	dll.InsertEnd(40)
	dll.InsertEnd(400)
	dll.InsertBeing(0)
	fmt.Println(dll.String())
	dll.reverse()
	fmt.Println(dll.String())
	dll.delHead()
	fmt.Println(dll.String())
	dll.delEnd()
	fmt.Println(dll.String())
	dll.delEnd()
	fmt.Println(dll.String())
}
