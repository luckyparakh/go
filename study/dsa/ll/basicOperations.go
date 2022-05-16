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
	ll.traverseRecur(ll.head)
	ll.addEnd(400)
	fmt.Println(ll.String())
	ll.delete()
	ll.delete()
	ll.delete()
	fmt.Println(ll.String())
	ll.delete()
	ll.addEnd(5)
	ll.addEnd(50)
	ll.deleteEnd()
	fmt.Println(ll.String())
	ll.deleteEnd()
	fmt.Println(ll.String())
	ll.addAtPosition(1, 10)
	ll.addAtPosition(1, 100)
	ll.addAtPosition(10, 10000) //will not add anything
	ll.addAtPosition(2, 1000)
	ll.addAtPosition(4, -1)
	fmt.Println(ll.String())
	fmt.Println(ll.search(10))
	fmt.Println(ll.search(20))
	fmt.Println(ll.searchRecur(10, ll.head, 1))
	fmt.Println(ll.searchRecur(20, ll.head, 1))
	fmt.Println(ll.searchRecur(100, ll.head, 1))
	fmt.Println(ll.searchRecurAnother(10, ll.head))
	fmt.Println(ll.searchRecurAnother(20, ll.head))
	fmt.Println(ll.searchRecurAnother(100, ll.head))
}

func (ll *linkedList) add(v int) {
	//Adding element at begining
	//TC:O(1)
	n := node{
		value: v,
	}
	n.next = ll.head
	ll.head = &n
}
func (ll *linkedList) addEnd(v int) {
	//Adding element at last
	n := node{
		value: v,
	}
	//Specail Case
	if ll.head == nil {
		n.next = ll.head // it is not needed but kept it for sake of understanding
		ll.head = &n
		return
	}
	temp1 := ll.head
	for temp1.next != nil { // Kind of while loop
		temp1 = temp1.next
	}
	n.next, temp1.next = temp1.next, &n
}

func (ll *linkedList) addAtPosition(position int, val int) {
	//think of about specail case like insert at position 1 and last position
	newNode := node{
		value: val,
	}
	counter := 1
	nodeCurr := ll.head
	var prevNode *node
	//Special Case
	if position == 1 {
		if ll.head == nil {
			ll.head = &newNode
			return
		}
		newNode.next, ll.head = ll.head, &newNode
		return
	}
	for nodeCurr != nil && counter < position {
		counter++
		prevNode = nodeCurr
		nodeCurr = nodeCurr.next
	}
	//Special Case
	if nodeCurr == nil && counter < position-1 {
		//When position is greater than len of LL
		return
	}
	newNode.next, prevNode.next = prevNode.next, &newNode
}

func (ll *linkedList) delete() {
	//Delete node at start
	//TC O(1)
	if ll.head != nil {
		ll.head = ll.head.next
	} else {
		fmt.Println("LL is empty.")
	}
}
func (ll *linkedList) deleteEnd() {
	//Delete node at start
	//TC O(n)
	if ll.head == nil {
		return
	}
	if ll.head.next == nil {
		ll.head = nil
		return
	}
	nodeCurr := ll.head
	for nodeCurr.next.next != nil {
		nodeCurr = nodeCurr.next
	}
	nodeCurr.next = nil
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
	// TC O(n), SC O(1)
	if n == nil {
		fmt.Println("")
		return
	}
	fmt.Printf("%d ", n.value)
	ll.traverseRecur(n.next)
}

func (ll *linkedList) search(val int) int {
	//TC O(n); SC O(1)
	//TC O(n)
	counter := 1
	//Specail case when LL is empty or when element is not present in LL
	for currNode := ll.head; currNode != nil; currNode = currNode.next {
		if currNode.value == val {
			return counter
		}
		counter++
	}
	return -1
}
func (ll *linkedList) searchRecur(val int, node *node, counter int) int {
	//TC O(n); SC O(n)
	// Can't apply binary search on sorted LL because it is not easy to find mid and to find lenght/mid of LL atleast one traversal is needed
	if node == nil {
		return -1
	}
	if node.value == val {
		return counter
	}
	return ll.searchRecur(val, node.next, counter+1)
}

func (ll *linkedList) searchRecurAnother(val int, node *node) int {
	if node == nil {
		return -1
	}
	if node.value == val {
		return 1
	}
	result := ll.searchRecurAnother(val, node.next)
	if result == -1 {
		return result
	} else {
		return result + 1
	}

}

func (ll *linkedList) String() string {
	//To string function
	//Use string builder because it is more efficient
	// TC O(n), SC O(n)
	var sb strings.Builder
	for node := ll.head; node != nil; node = node.next {
		sb.WriteString(fmt.Sprintf("%d ", node.value))
	}
	return sb.String()
}
