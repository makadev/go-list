package go_list

// SingleLinkedList is a generic type representing a single linked list.
type SingleLinkedList[T any] struct {
	// The first element of the list.
	Head *SingleLinkedListElement[T]
	// The last element of the list.
	Tail *SingleLinkedListElement[T]
	// The number of elements in the list.
	Count uint64
}

// SingleLinkedListElement is an element of SingleLinkedList.
type SingleLinkedListElement[T any] struct {
	// The next element in the list.
	Next *SingleLinkedListElement[T]
	// The value of the element.
	Value T
}

// Init initializes or clears list l.
func (l *SingleLinkedList[T]) Init() {
	l.Head = nil
	l.Tail = nil
	l.Count = 0
}

// PushBack inserts a new SingleLinkedListElement with given value at the back of list l.
func (l *SingleLinkedList[T]) PushBack(value T) {
	node := &SingleLinkedListElement[T]{Next: nil, Value: value}
	if l.Head == nil {
		l.Head = node
	} else {
		l.Tail.Next = node
	}
	l.Tail = node
	l.Count++
}

// PushFront inserts a new SingleLinkedListElement with given value at the front of list l.
func (l *SingleLinkedList[T]) PushFront(value T) {
	node := &SingleLinkedListElement[T]{Next: l.Head, Value: value}
	l.Head = node
	if l.Tail == nil {
		l.Tail = node
	}
	l.Count++
}

// PopFront removes the first element from list l.
func (l *SingleLinkedList[T]) PopFront() {
	if l.Head == nil {
		return
	}
	var oldHead = l.Head
	l.Head = l.Head.Next
	if l.Head == nil {
		l.Tail = nil
	}
	// explicit unref
	oldHead.Next = nil
	l.Count--
}

// ToArray returns an array of all elements in list l.
func (l *SingleLinkedList[T]) ToArray() []T {
	var arr []T = make([]T, 0, l.Count)
	for node := l.Head; node != nil; node = node.Next {
		arr = append(arr, node.Value)
	}
	return arr
}

// SingleLinkedListIterator is an iterator for SingleLinkedList.
//
// Example:
//
//	iterator := &SingleLinkedListIterator[int]{}
//	iterator.Init(list)
//	for iterator.Next() {
//		fmt.Println(iterator.Value())
//	}
type SingleLinkedListIterator[T any] struct {
	List    *SingleLinkedList[T]
	Current *SingleLinkedListElement[T]
}

// Init initializes or resets the iterator with the given list.
func (i *SingleLinkedListIterator[T]) Init(list *SingleLinkedList[T]) {
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
func (i *SingleLinkedListIterator[T]) Next() bool {
	if i.Current != nil {
		i.Current = i.Current.Next
	} else {
		i.Current = i.List.Head
	}
	return i.Current != nil
}

// Value returns the value of the current element.
//
// Do not call Value() if Next() returned false.
func (i *SingleLinkedListIterator[T]) Value() T {
	return i.Current.Value
}
