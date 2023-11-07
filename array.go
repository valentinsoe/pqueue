package pqueue

import (
	"fmt"
)

// Implementation of priority queue using array
type array struct {
	data []Item
}

// Creates new array-based priority queue
func NewArrayPQueue() *array {
	return &array{
		data: make([]Item, 0),
	}
}

func (apq *array) Enqueue(item Item) {
	apq.data = append(apq.data, item)
}

func (apq *array) Dequeue() (item Item, ok bool) {
	if index, ok := apq.peek(); ok {
		item = apq.data[index]

		apq.remove(index)

		return item, true
	}

	return Item{}, false
}

func (apq *array) Peek() (item Item, ok bool) {
	if index, ok := apq.peek(); ok {
		return apq.data[index], true
	}
	return Item{}, false

}

func (apq *array) Remove(item Item) bool {
	if index, ok := apq.find(item); ok {
		apq.remove(index)
		return true
	}
	return false
}

func (apq *array) ChangePriority(item Item, priority int) bool {
	if index, ok := apq.find(item); ok {
		apq.data[index].Priority = priority
		return true
	}
	return false
}

func (apq *array) peek() (index int, ok bool) {
	if len(apq.data) == 0 {
		return 0, false
	}

	//we assume that priorities are positive
	maxPriority := -1
	maxItemIndex := -1
	for i, item := range apq.data {
		if item.Priority > maxPriority || (item.Priority == maxPriority && item.Value > apq.data[maxItemIndex].Value) {
			maxPriority = item.Priority
			maxItemIndex = i
		}
	}
	return maxItemIndex, true
}

func (apq *array) remove(index int) {
	//We need to move all elements after index one position left
	for i := index + 1; i < len(apq.data); i++ {
		apq.data[i-1] = apq.data[i]
	}
	apq.data = apq.data[:len(apq.data)-1]
}

func (apq *array) find(item Item) (index int, ok bool) {
	for i, v := range apq.data {
		if v.Value == item.Value && v.Priority == item.Priority {
			return i, true
		}
	}
	return 0, false
}

func (apq *array) Print() {
	for i, v := range apq.data {
		fmt.Printf("[%d]: %d\r\n", i, v)
	}
}
