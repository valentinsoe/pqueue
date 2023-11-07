// This package show several implementations of priority queue for educational purposes
package pqueue

// Single item of priority queue
type Item struct {
	//Value of item
	//Can be generic, but left as int for simplicity
	Value int
	//Priority of item
	Priority int
}

func (item Item) less(another Item) bool {
	return item.Priority < another.Priority ||
		(item.Priority == another.Priority && item.Value < another.Value)
}

// Interface of priority queue
type PQueue interface {
	//Adds element to queue
	Enqueue(item Item)
	//Gets and removes element with highest priority from the queue
	//Returns false if there are no elements
	Dequeue() (item Item, ok bool)
	//Gets(and not removes) element with highest priority from the queue
	//Returns false if there are no elements
	Peek() (item Item, ok bool)
	//Removes element with same value and priority from the queue
	//Returns true if the element was found
	//Returns false if the element wasn't found
	Remove(item Item) bool
	//Changes element priorotity for the element with same value and priority
	//Returns true if the element was found
	//Returns false if the element wasn't found
	ChangePriority(item Item, priority int) bool
	//Print out for debug purposes
	Print()
}
