package pqueue

import (
	"fmt"
)

// Implementation of priority queue using sorted list
type sortedList struct {
	first *sortedListItem
}

type sortedListItem struct {
	Item
	next *sortedListItem
}

// Creates new list-based priority queue
func NewSortedListPQueue() *sortedList {
	return &sortedList{}
}

func (slpq *sortedList) Enqueue(item Item) {
	if slpq.first == nil {
		slpq.first = &sortedListItem{
			Item: item,
		}
		return
	}
	if slpq.first.Priority < item.Priority ||
		(slpq.first.Priority == item.Priority && slpq.first.Value < item.Value) {
		slpq.first = &sortedListItem{
			Item: item,
			next: slpq.first,
		}
		return
	}

	for listItem := slpq.first; listItem != nil; listItem = listItem.next {
		if listItem.next == nil {
			listItem.next = &sortedListItem{
				Item: item,
			}
			return
		}
		if item.Priority > listItem.next.Priority ||
			(item.Priority == listItem.next.Priority && item.Value > listItem.next.Value) {
			listItem.next = &sortedListItem{
				Item: item,
				next: listItem.next,
			}
			return
		}
	}
}

func (slpq *sortedList) Dequeue() (item Item, ok bool) {
	if slpq.first == nil {
		return Item{}, false
	}

	item = slpq.first.Item
	slpq.first = slpq.first.next

	return item, true
}

func (slpq *sortedList) Peek() (item Item, ok bool) {
	if slpq.first != nil {
		return slpq.first.Item, true
	}
	return Item{}, false

}

func (slpq *sortedList) Remove(item Item) bool {
	for prevListItem, listItem := slpq.first, slpq.first; listItem != nil; prevListItem, listItem = listItem, listItem.next {
		if listItem.Value == item.Value && listItem.Priority == item.Priority {
			if listItem == slpq.first {
				slpq.first = listItem.next
			} else {
				prevListItem.next = listItem.next
			}
			return true
		}
	}
	return false
}

func (slpq *sortedList) ChangePriority(item Item, priority int) bool {
	if listItem, ok := slpq.find(item); ok {
		listItem.Priority = priority
		return true
	}
	return false
}

func (slpq *sortedList) find(item Item) (listItem *sortedListItem, ok bool) {
	for listItem := slpq.first; listItem != nil; listItem = listItem.next {
		if listItem.Value == item.Value && listItem.Priority == item.Priority {
			return listItem, true
		}
	}
	return nil, false
}

func (slpq *sortedList) Print() {
	for i, listItem := 0, slpq.first; listItem != nil; i, listItem = i+1, listItem.next {
		fmt.Printf("[%d]: %d\r\n", i, listItem.Item)
	}
}
