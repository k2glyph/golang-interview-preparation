package main

import "fmt"

// node struct
type node struct {
	data int
	next *node
}

//linkedlist struct to hold information of linkedlist
type linkedlist struct {
	head   *node
	length int
}

//additem to the linkedlist
func (ll *linkedlist) addItem(items ...*node) {
	for _, item := range items {
		temp := ll.head
		ll.head = item
		ll.head.next = temp
		ll.length++
	}

}

// deleteitem from the linkedlist
func (ll *linkedlist) deleteItemWithValue(value int) {
	if ll.length == 0 {
		return
	}
	temp := ll.head
	if temp.data == value {
		ll.head = ll.head.next
		ll.length--
		return
	}
	for temp.next.data != value {
		if temp.next.next == nil {
			return
		}
		temp = temp.next
	}
	temp.next = temp.next.next

	ll.length--
}

// print the list of item in linkedlist
func (ll *linkedlist) printItem() {
	temp := ll.head
	length := ll.length
	for length != 0 {
		fmt.Printf("%d ", temp.data)
		temp = temp.next
		length--
	}
	fmt.Printf("\n")
}
func main() {
	mylist := linkedlist{}
	nodes := []*node{{data: 10}, {data: 11}, {data: 12}, {data: 13}, {data: 14}, {data: 15}, {data: 16}}
	mylist.addItem(nodes...)
	mylist.printItem()
	mylist.deleteItemWithValue(12)
	mylist.printItem()
	mylist.deleteItemWithValue(13)
	mylist.printItem()
	mylist.deleteItemWithValue(16)
	mylist.printItem()
	mylist.deleteItemWithValue(10)
	mylist.printItem()
}
