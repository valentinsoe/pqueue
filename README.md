# pqueue
Several variants of priority queue implementations.

array - storing data in array, and lookup for highest priority for each Dequeue.
  Fast Enqueue, very slow Dequeue
  
sorted_list - storing data in linked list. On each Enqueue traverse list to find correct place for the item.
  Fast Dequeue, very slow Enqueue

heap - storing data in complete binary tree, represented as array. 
  On each Enqueue - adds item as last leaf and shifts it up according to priority.
  On each Dequeue - replaces root element with last leaf and shifts it down according to priority.
  Marginally slower Enqueue than array, marginally slower Dequeue than sorted_list
