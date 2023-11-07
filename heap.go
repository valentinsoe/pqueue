package pqueue

import (
	"fmt"
)

// Implementation of priority queue using heap
type heap struct {
	data []Item
}

// Creates new heap-based priority queue
func NewHeapPQueue() *heap {
	return &heap{}
}

// Function to return the index of the parent item of a given item
func parent(index int) int {
	return (index - 1) / 2
}

// Function to return the index of the right child of the given item
func rightChild(index int) int {
	return ((2 * index) + 2)
}

// Function to return the index of the left child of the given item
func leftChild(index int) int {
	return ((2 * index) + 1)
}

// Function to swap two items by index
func (hpq *heap) swap(i, j int) {
	item := hpq.data[i]
	hpq.data[i] = hpq.data[j]
	hpq.data[j] = item
}

// Function to shift up the node in order to maintain the heap property
func (hpq *heap) shiftUp(index int) {
	for index > 0 && hpq.data[parent(index)].less(hpq.data[index]) {
		// Swap parent and current node
		hpq.swap(parent(index), index)
		// Update index to parent of i
		index = parent(index)
	}
}

// Function to shift down the node in order to maintain the heap property
func (hpq *heap) shiftDown(index int) {
	maxIndex := index

	// Left Child
	left := leftChild(index)
	if left < len(hpq.data) && hpq.data[maxIndex].less(hpq.data[left]) {
		maxIndex = left
	}

	// Right Child
	right := rightChild(index)
	if right < len(hpq.data) && hpq.data[maxIndex].less(hpq.data[right]) {
		maxIndex = right
	}

	// If i not same as maxIndex
	if index != maxIndex {
		hpq.swap(index, maxIndex)
		hpq.shiftDown(maxIndex)
	}
}

func (hpq *heap) Enqueue(item Item) {
	hpq.data = append(hpq.data, item)
	hpq.shiftUp(len(hpq.data) - 1)
}

func (hpq *heap) Dequeue() (item Item, ok bool) {
	if len(hpq.data) == 0 {
		return Item{}, false
	}

	result := hpq.data[0]

	// Replace the value at the root with the last leaf
	hpq.data[0] = hpq.data[len(hpq.data)-1]
	hpq.data = hpq.data[:len(hpq.data)-1]

	// Shift down the replaced element to maintain the heap property
	hpq.shiftDown(0)

	return result, true
}

func (hpq *heap) Peek() (item Item, ok bool) {
	if len(hpq.data) == 0 {
		return Item{}, false
	}

	return hpq.data[0], true
}

func (hpq *heap) Remove(item Item) bool {
	if index, ok := hpq.find(item); ok {
		hpq.data[index].Priority = hpq.data[0].Priority + 1
		hpq.shiftUp(index)
		hpq.Dequeue()
		return true
	}
	return false
}

func (hpq *heap) ChangePriority(item Item, priority int) bool {
	if index, ok := hpq.find(item); ok {
		oldPriority := hpq.data[index].Priority
		hpq.data[index].Priority = priority
		if priority > oldPriority {
			hpq.shiftUp(index)
		} else {
			hpq.shiftDown(index)
		}
		return true
	}
	return false
}

func (hpq *heap) Print() {
	for i, v := range hpq.data {
		fmt.Printf("[%d]: %d\r\n", i, v)
	}
}

//looks up for the item in the heap(takes O(n))
func (hpq *heap) find(item Item) (index int, ok bool) {
	for i, v := range hpq.data {
		if v.Value == item.Value && v.Priority == item.Priority {
			return i, true
		}
	}
	return 0, false
}
