package utils

import "fmt"

type Item struct {
	Data int
	Next *Item
}

var (
	Head Item
	Tail Item
)

func check(it Item) bool {
	var slow Item = it
	var fast Item = it
	var fast2 Item = it
	for {
		if &slow == nil || fast2.Next == nil || fast2.Next.Next == nil {
			break;
		}
		fast = *fast2.Next
		fast2 = *fast.Next
		if slow == fast || slow == fast2 {
			return true
		}
		slow = *slow.Next
	}
	return false
}

func CheckLoopForLinkList() {
	var item Item
	var item1 Item
	var item2 Item
	var item3 Item
	var item4 Item
	var item5 Item
	var item6 Item
	// Looped link list
	item.Next = &item1
	item1.Next = &item2
	item2.Next = &item3
	item3.Next = &item4
	item4.Next = &item5
	item5.Next = &item6
	item6.Next = nil
	fmt.Println(check(item))
}
