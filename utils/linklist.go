package utils

import (
	"fmt"
)

type Item struct {
	Data int
	Next *Item
}

// Insert to head of list
func (it *Item) insert(data int) *Item {
	return &Item{Data: data, Next: it}
}

func (it *Item) merge(newItem *Item) *Item {
	current, next, next2 := it, it, it
	_, _ = next, next2
	for ; current.Next != nil; {
		current = current.Next
	}
	current.Next = newItem
	return it
}

// Insert to tail of list
func (it *Item) add(data int) *Item {
	current := it
	count := 0
	for {
		if current.Next == nil {
			current.Next = &Item{Data: data}
			break
		}
		count++
		current = current.Next
	}
	return it
}

func (it *Item) printList() {
	current := it
	count := 0
	for current != nil {
		fmt.Printf("[%d]: %d\n", count, current.Data)
		count++
		current = current.Next
	}
}

func (it *Item) printReverseList(count ...int) {
	cnt := 0
	if count != nil {
		cnt = count[0] + 1
	}
	if it.Next != nil {
		it.Next.printReverseList(cnt)
	}
	fmt.Printf("[%#v]: %d\n", cnt, it.Data)
}

var (
	Head Item
	Tail Item
)

func check(it Item) bool {
	var slow Item = it
	var fast Item = it
	var fast2 Item = it
	count := 0
	for {
		if &slow == nil || fast2.Next == nil || fast2.Next.Next == nil {
			break
		}
		fast = *fast2.Next
		fast2 = *fast.Next
		count++
		println(count, slow.Data, fast.Data, fast2.Data)
		if slow == fast || slow == fast2 {
			return true
		}
		slow = *slow.Next
	}
	return false
}

type a1 struct {
	x int
}

type a2 struct {
	a1
	y int
}

func (a a1) tst() {
	println("call", a.x)
}

func (a *a1) tst2() {
	println("call", a.x)
}

func CheckLoopForLinkList() {
	var a, b *a1
	var c a1
	a = &a1{1}
	b = a
	b.x = 1
	a.tst()
	b.tst()
	c.tst()
	(*a).tst()
	(*b).tst()
	(&c).tst()

	var item,
		item1,
		item2,
		item3,
		item4,
		item5,
		item6,
		item7,
		item8,
		item9,
		item10 Item

	item  = Item{Data:0}
	item1 = Item{Data:1}
	item2 = Item{Data:2}
	item3 = Item{Data:3}
	item4 = Item{Data:4}
	item5 = Item{Data:5}
	item6 = Item{Data:6}
	item7 = Item{Data:7}
	item8 = Item{Data:8}
	item9 = Item{Data:9}
	item10 = Item{Data:10}

	item.Next = &item1
	item1.Next = &item2
	item2.Next = &item3
	item3.Next = &item4
	item4.Next = &item10
	item5.Next = &item6
	item6.Next = &item7
	item7.Next = &item8
	item8.Next = &item9
	item9.Next = &item10
	item10.Next = &item6
	fmt.Println(check(item))

	var it *Item
	list1 := it.
		insert(10).
		insert(22).
		insert(30).
		insert(44).
		insert(50).
		insert(66).
		insert(70).
		insert(88)
	println("Inserted List: ======")
	list1.printList()
	println("Reversed Inserted List: ======")
	list1.printReverseList()

	var it2 Item
	list2 := it2.
		add(10).
		add(22).
		add(30).
		add(44).
		add(50).
		add(66).
		add(70).
		add(88)
	println("Append List: ======")
	list2.printList()
	println("Reversed Append List: ======")
	list2.printReverseList()

	list1.merge(list2)
	if check(*list1) {
		println("Linked List loop detected =====")
	}
	println("Merged List: ======")
	list1.printList()
	
	list1.merge(list2)
	if check(*list1) {
		println("Linked List loop detected =====")
	} else {
		list1.printList()
	}
}
