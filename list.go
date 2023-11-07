package pqueue

import (
	"fmt"
)

// Implementation of priority queue using list
type list struct {
	first *listItem
	last  *listItem
}

type listItem struct {
	Item
	next *listItem
	prev *listItem
}

// Creates new list-based priority queue
func NewListPQueue() *list {
	return &list{}
}

func (lpq *list) Enqueue(item Item) {
	if lpq.first == nil {
		lpq.first = &listItem{
			Item: item,
		}
		lpq.last = lpq.first
	} else {
		newListItem := &listItem{
			Item: item,
			prev: lpq.last,
		}
		lpq.last.next = newListItem
		lpq.last = newListItem
	}
}

func (lpq *list) Dequeue() (item Item, ok bool) {
	if listItem, ok := lpq.peek(); ok {
		lpq.remove(listItem)

		return listItem.Item, true
	}

	return Item{}, false
}

func (lpq *list) Peek() (item Item, ok bool) {
	if listItem, ok := lpq.peek(); ok {
		return listItem.Item, true
	}
	return Item{}, false

}

func (lpq *list) Remove(item Item) bool {
	if listItem, ok := lpq.find(item); ok {
		lpq.remove(listItem)
		return true
	}
	return false
}

func (lpq *list) ChangePriority(item Item, priority int) bool {
	if listItem, ok := lpq.find(item); ok {
		listItem.Priority = priority
		return true
	}
	return false
}

func (lpq *list) peek() (item *listItem, ok bool) {
	if lpq.first == nil {
		return nil, false
	}

	//we assume that priorities are positive
	maxPriority := -1
	var maxItem *listItem = nil
	for item := lpq.first; item != nil; item = item.next {
		if item.Priority > maxPriority || (item.Priority == maxPriority && item.Value > maxItem.Value) {
			maxPriority = item.Priority
			maxItem = item
		}
	}
	return maxItem, true
}

func (lpq *list) remove(item *listItem) {
	if lpq.first == item {
		lpq.first = item.next
	}
	if lpq.last == item {
		lpq.last = item.prev
	}
	if item.next != nil {
		item.next.prev = item.prev
	}
	if item.prev != nil {
		item.prev.next = item.next
	}
}

func (lpq *list) find(item Item) (listItem *listItem, ok bool) {
	for listItem := lpq.first; listItem != nil; listItem = listItem.next {
		if listItem.Value == item.Value && listItem.Priority == item.Priority {
			return listItem, true
		}
	}
	return nil, false
}

func (lpq *list) Print() {
	for i, listItem := 0, lpq.first; listItem != nil; i, listItem = i+1, listItem.next {
		fmt.Printf("[%d]: %d\r\n", i, listItem.Item)
	}
}
