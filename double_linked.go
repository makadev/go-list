package go_list

// DoubleLinkedList is a generic type representing a doubly linked list.
type DoubleLinkedList[T any] struct {
	Head  *DoubleLinkedListElement[T]
	Tail  *DoubleLinkedListElement[T]
	Count uint64
}

// DoubleLinkedListElement is an element of DoubleLinkedList.
type DoubleLinkedListElement[T any] struct {
	Next  *DoubleLinkedListElement[T]
	Prev  *DoubleLinkedListElement[T]
	Value T
}

// Init initializes or clears list l.
func (l *DoubleLinkedList[T]) Init() {
	l.Head = nil
	l.Tail = nil
	l.Count = 0
}

// PushBack inserts a new DoubleLinkedListElement with given value at the back of list l.
func (l *DoubleLinkedList[T]) PushBack(value T) {
	node := &DoubleLinkedListElement[T]{Next: nil, Prev: l.Tail, Value: value}
	if l.Head == nil {
		l.Head = node
	} else {
		l.Tail.Next = node
	}
	l.Tail = node
	l.Count++
}

// PushFront inserts a new DoubleLinkedListElement with given value at the front of list l.
func (l *DoubleLinkedList[T]) PushFront(value T) {
	node := &DoubleLinkedListElement[T]{Next: l.Head, Prev: nil, Value: value}
	l.Head = node
	if l.Tail == nil {
		l.Tail = node
	}
	l.Count++
}

// PopFront removes the first element from list l.
func (l *DoubleLinkedList[T]) PopFront() {
	if l.Head == nil {
		return
	}
	var oldHead = l.Head
	l.Head = l.Head.Next
	if l.Head == nil {
		l.Tail = nil
	} else {
		l.Head.Prev = nil
	}
	// explicitly unref
	oldHead.Next = nil
	oldHead.Prev = nil
	l.Count--
}

// PopBack removes the last element from list l.
func (l *DoubleLinkedList[T]) PopBack() {
	if l.Tail == nil {
		return
	}
	var oldTail = l.Tail
	l.Tail = l.Tail.Prev
	if l.Tail == nil {
		l.Head = nil
	} else {
		l.Tail.Next = nil
	}
	// explicitly unref
	oldTail.Next = nil
	oldTail.Prev = nil
	l.Count--
}

// ToArray returns an array of all elements in list l.
func (l *DoubleLinkedList[T]) ToArray() []T {
	var arr []T = make([]T, 0, l.Count)
	for node := l.Head; node != nil; node = node.Next {
		arr = append(arr, node.Value)
	}
	return arr
}

// DoubleLinkedListIterator is an iterator for DoubleLinkedList.
//
// Forward Example:
//
//	iterator := &DoubleLinkedListIterator[int]{}
//	iterator.Init(list)
//	for iterator.Next() {
//		fmt.Println(iterator.Value())
//	}
//
// Backward Example:
//
//	iterator := &DoubleLinkedListIterator[int]{}
//	iterator.Init(list)
//	for iterator.Prev() {
//		fmt.Println(iterator.Value())
//	}
type DoubleLinkedListIterator[T any] struct {
	List    *DoubleLinkedList[T]
	Current *DoubleLinkedListElement[T]
}

// Init initializes or resets the iterator with the given list.
func (i *DoubleLinkedListIterator[T]) Init(list *DoubleLinkedList[T]) {
	i.List = list
	i.Current = nil
}

// Next moves the iterator to the next element and returns true if there was a next element.
//
// Do not call Next() after it returned false since it will reset the iterator to the front of the list.
// Example:
//
//	for iterator.Next() {
//		fmt.Println(iterator.Value())
//	}
//	iterator.Current = nil // explicitly reset the iterator
//
//	for iterator.Next() {
//		fmt.Println(iterator.Value())
//	}
func (i *DoubleLinkedListIterator[T]) Next() bool {
	if i.Current != nil {
		i.Current = i.Current.Next
	} else {
		i.Current = i.List.Head
	}
	return i.Current != nil
}

// Prev moves the iterator to the previous element and returns true if there was a previous element.
//
// Do not call Prev() after it returned false since it will reset the iterator to the back of the list.
// Example:
//
//	for iterator.Prev() {
//		fmt.Println(iterator.Value())
//	}
//	iterator.Current = nil // explicitly reset the iterator
//
//	for iterator.Prev() {
//		fmt.Println(iterator.Value())
//	}
func (i *DoubleLinkedListIterator[T]) Prev() bool {
	if i.Current != nil {
		i.Current = i.Current.Prev
	} else {
		i.Current = i.List.Tail
	}
	return i.Current != nil
}

// Value returns the value of the current element.
//
// Do not call Value() if Next() or Prev() returned false.
func (i *DoubleLinkedListIterator[T]) Value() T {
	return i.Current.Value
}
