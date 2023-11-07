package pqueue

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

type queueDef struct {
	name      string
	queueFunc func() PQueue
}

var queues = []queueDef{
	{"array", func() PQueue { return NewArrayPQueue() }},
	//{"list", func() PQueue { return NewListPQueue() }},
	{"sortedList", func() PQueue { return NewSortedListPQueue() }},
	{"heap", func() PQueue { return NewHeapPQueue() }},
}

var benchmarkSize = []int{
	1000,
	100000,
}

func TestQueue(t *testing.T) {
	for _, queueDef := range queues {
		testCases := []struct {
			name          string
			itemsToAdd    []Item
			itemsToRemove []Item
			itemsToGet    []Item
		}{
			{
				name:       fmt.Sprintf("%s one item", queueDef.name),
				itemsToAdd: []Item{Item{1, 2}},
				itemsToGet: []Item{Item{1, 2}},
			},
			{
				name:       fmt.Sprintf("%s two direct order", queueDef.name),
				itemsToAdd: []Item{Item{1, 2}, Item{1, 1}},
				itemsToGet: []Item{Item{1, 2}, Item{1, 1}},
			},
			{
				name:       fmt.Sprintf("%s two wrong order", queueDef.name),
				itemsToAdd: []Item{Item{1, 2}, Item{1, 3}},
				itemsToGet: []Item{Item{1, 3}, Item{1, 2}},
			},
			{
				name:       fmt.Sprintf("%s many items same priority", queueDef.name),
				itemsToAdd: []Item{Item{1, 2}, Item{2, 2}, Item{3, 2}, Item{100, 2}},
				itemsToGet: []Item{Item{100, 2}, Item{3, 2}, Item{2, 2}, Item{1, 2}},
			},
			{
				name:          fmt.Sprintf("%s many items remove middle", queueDef.name),
				itemsToAdd:    []Item{Item{1, 2}, Item{2, 2}, Item{3, 2}, Item{100, 2}},
				itemsToRemove: []Item{Item{2, 2}},
				itemsToGet:    []Item{Item{100, 2}, Item{3, 2}, Item{1, 2}},
			},
			{
				name:          fmt.Sprintf("%s many items remove high", queueDef.name),
				itemsToAdd:    []Item{Item{1, 2}, Item{2, 2}, Item{3, 2}, Item{100, 2}},
				itemsToRemove: []Item{Item{100, 2}},
				itemsToGet:    []Item{Item{3, 2}, Item{2, 2}, Item{1, 2}},
			},
			{
				name:          fmt.Sprintf("%s many items remove low", queueDef.name),
				itemsToAdd:    []Item{Item{1, 2}, Item{2, 2}, Item{3, 2}, Item{100, 2}},
				itemsToRemove: []Item{Item{1, 2}},
				itemsToGet:    []Item{Item{100, 2}, Item{3, 2}, Item{2, 2}},
			},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				queue := queueDef.queueFunc()
				for _, v := range tc.itemsToAdd {
					queue.Enqueue(v)
				}

				for _, v := range tc.itemsToRemove {
					ok := queue.Remove(v)
					assert.True(t, ok, "Item not removed")
				}

				for _, expected := range tc.itemsToGet {
					item, ok := queue.Dequeue()
					assert.True(t, ok, "Item not dequeued")
					assert.Equal(t, expected, item, "Item is not the same")
				}
				_, ok := queue.Dequeue()
				assert.False(t, ok, "Item dequeued")
			})
		}
	}
}

func BenchmarkEnqueue(b *testing.B) {
	for _, queueDef := range queues {
		for _, size := range benchmarkSize {
			queue := queueDef.queueFunc()
			b.Run(fmt.Sprintf("%s %d", queueDef.name, size), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					for i := 0; i < size; i++ {
						queue.Enqueue(Item{rand.Int(), rand.Int()})
					}
				}
			})
		}
	}
}

/*func BenchmarkPeek(b *testing.B) {
	for _, queueDef := range queues {
		for _, size := range benchmarkSize {
			queue := queueDef.queueFunc()
			for i := 0; i < size; i++ {
				queue.Enqueue(Item{rand.Int(), rand.Int()})
			}
			b.Run(fmt.Sprintf("%s %d", queueDef.name, size), func(b *testing.B) {

				for i := 0; i < b.N; i++ {
					for i := 0; i < size; i++ {
						queue.Peek()
					}
				}
			})
		}
	}
}*/

func BenchmarkDequeue(b *testing.B) {
	for _, queueDef := range queues {
		for _, size := range benchmarkSize {
			queue := queueDef.queueFunc()
			for i := 0; i < size; i++ {
				queue.Enqueue(Item{rand.Int(), rand.Int()})
			}
			b.Run(fmt.Sprintf("%s %d", queueDef.name, size), func(b *testing.B) {

				for i := 0; i < b.N; i++ {
					for i := 0; i < size; i++ {
						queue.Dequeue()
					}
				}
			})
		}
	}
}
