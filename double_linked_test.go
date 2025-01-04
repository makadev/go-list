package go_list

import (
	"math/rand"
	"testing"

	"github.com/go-test/deep"
)

func checkDListNotEmpty(t *testing.T, list *DoubleLinkedList[int], expectedCount uint64) {
	if list.Head == nil {
		t.Errorf("list.Head = nil; want !nil")
	}
	if list.Tail == nil {
		t.Errorf("list.Tail = nil; want !nil")
	}
	if list.Count != expectedCount {
		t.Errorf("list.Count = %v; want %v", list.Count, expectedCount)
	}
	// check that head / tail element are first / last
	if list.Head.Prev != nil {
		t.Errorf("list.Head.Prev = %v; want nil", list.Head.Prev)
	}
	if list.Tail.Next != nil {
		t.Errorf("list.Tail.Next = %v; want nil", list.Tail.Next)
	}
}

func checkDListIsEmpty(t *testing.T, list *DoubleLinkedList[int]) {
	list2 := &DoubleLinkedList[int]{
		Head:  nil,
		Tail:  nil,
		Count: 0,
	}
	if diff := deep.Equal(list, list2); diff != nil {
		t.Error(diff)
	}
}

func checkDListEquals(t *testing.T, list *DoubleLinkedList[int], arr []int) {
	if diff := deep.Equal(list.ToArray(), arr); diff != nil {
		t.Error(diff)
	}
}

func TestDoubleLinkedListInit(t *testing.T) {
	t.Run("Init", func(t *testing.T) {
		list := &DoubleLinkedList[int]{}
		list.Init()

		checkDListIsEmpty(t, list)
	})
}

func TestDoubleLinkedListAppend(t *testing.T) {
	t.Run("Append", func(t *testing.T) {
		list := &DoubleLinkedList[int]{}
		list.Init()

		checkDListIsEmpty(t, list)

		list.PushBack(1)

		checkDListNotEmpty(t, list, 1)
		checkDListEquals(t, list, []int{1})

		list.PushBack(2)

		checkDListNotEmpty(t, list, 2)
		checkDListEquals(t, list, []int{1, 2})
	})
}

func TestDoubleLinkedListPrepend(t *testing.T) {
	t.Run("Prepend", func(t *testing.T) {
		list := &DoubleLinkedList[int]{}
		list.Init()

		checkDListIsEmpty(t, list)

		list.PushFront(1)

		checkDListNotEmpty(t, list, 1)
		checkDListEquals(t, list, []int{1})

		list.PushFront(2)

		checkDListNotEmpty(t, list, 2)
		checkDListEquals(t, list, []int{2, 1})
	})
}

func TestDoubleLinkedListRemoveFront(t *testing.T) {
	t.Run("RemoveFront", func(t *testing.T) {
		list := &DoubleLinkedList[int]{}
		list.Init()

		list.PushBack(1)
		list.PushBack(2)
		list.PushBack(3)

		checkDListNotEmpty(t, list, 3)
		checkDListEquals(t, list, []int{1, 2, 3})

		list.PopFront()

		checkDListNotEmpty(t, list, 2)
		checkDListEquals(t, list, []int{2, 3})

		list.PopFront()

		checkDListNotEmpty(t, list, 1)
		checkDListEquals(t, list, []int{3})

		list.PopFront()

		checkDListIsEmpty(t, list)

		list.PopFront()

		checkDListIsEmpty(t, list)
	})
}

func TestDoubleLinkedListRemoveBack(t *testing.T) {
	t.Run("RemoveBack", func(t *testing.T) {
		list := &DoubleLinkedList[int]{}
		list.Init()

		list.PushBack(1)
		list.PushBack(2)
		list.PushBack(3)

		checkDListNotEmpty(t, list, 3)
		checkDListEquals(t, list, []int{1, 2, 3})

		list.PopBack()
		t.Log("check")
		checkDListNotEmpty(t, list, 2)
		checkDListEquals(t, list, []int{1, 2})

		list.PopBack()

		t.Log("check")
		checkDListNotEmpty(t, list, 1)
		checkDListEquals(t, list, []int{1})

		list.PopBack()

		checkDListIsEmpty(t, list)

		list.PopBack()

		checkDListIsEmpty(t, list)
	})
}

func TestDoubleLinkedListIteratorForward(t *testing.T) {
	t.Run("IteratorForward", func(t *testing.T) {
		list := &DoubleLinkedList[int]{}
		list.Init()

		list.PushBack(1)
		list.PushBack(2)
		list.PushBack(3)

		checkDListNotEmpty(t, list, 3)
		checkDListEquals(t, list, []int{1, 2, 3})

		iter := DoubleLinkedListIterator[int]{}
		iter.Init(list)

		var arr []int = make([]int, 0, list.Count)
		for iter.Next() {
			arr = append(arr, iter.Value())
		}

		if diff := deep.Equal(arr, []int{1, 2, 3}); diff != nil {
			t.Error(diff)
		}
	})
}

func TestDoubleLinkedListIteratorBackward(t *testing.T) {
	t.Run("IteratorBackward", func(t *testing.T) {
		list := &DoubleLinkedList[int]{}
		list.Init()

		list.PushBack(1)
		list.PushBack(2)
		list.PushBack(3)

		checkDListNotEmpty(t, list, 3)
		checkDListEquals(t, list, []int{1, 2, 3})

		iter := DoubleLinkedListIterator[int]{}
		iter.Init(list)

		var arr []int = make([]int, 0, list.Count)
		for iter.Prev() {
			arr = append(arr, iter.Value())
		}

		if diff := deep.Equal(arr, []int{3, 2, 1}); diff != nil {
			t.Error(diff)
		}
	})
}

func BenchmarkDoubleLinkedListAddDelete(b *testing.B) {
	list := &DoubleLinkedList[int]{}
	list.Init()
	for i := 0; i < b.N; i++ {
		list.PushBack(i)
	}
	for i := 0; i < b.N; i++ {
		list.PopBack()
	}
}

func BenchmarkDoubleLinkedListMixedAddDelete(b *testing.B) {
	list := &DoubleLinkedList[int]{}
	list.Init()
	for i := 0; i < b.N; i++ {
		if i%2 == 0 {
			list.PushBack(i)
		} else {
			list.PushFront(i)
		}
	}
	for i := 0; i < b.N; i++ {
		if rand.Intn(2) == 0 {
			list.PopBack()
		} else {
			list.PopFront()
		}
	}
}
