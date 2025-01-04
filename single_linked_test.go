package go_list

import (
	"testing"

	"github.com/go-test/deep"
)

func checkListNotEmpty(t *testing.T, list *SingleLinkedList[int], expectedCount uint64) {
	if list.Head == nil {
		t.Errorf("list.Head = nil; want !nil")
	}
	if list.Tail == nil {
		t.Errorf("list.Tail = nil; want !nil")
	}
	if list.Count != expectedCount {
		t.Errorf("list.Count = %v; want %v", list.Count, expectedCount)
	}
	// check that tail element is last
	if list.Tail.Next != nil {
		t.Errorf("list.Tail.Next = %v; want nil", list.Tail.Next)
	}
}

func checkListIsEmpty(t *testing.T, list *SingleLinkedList[int]) {
	list2 := &SingleLinkedList[int]{
		Head:  nil,
		Tail:  nil,
		Count: 0,
	}
	if diff := deep.Equal(list, list2); diff != nil {
		t.Error(diff)
	}
}

func checkListEquals(t *testing.T, list *SingleLinkedList[int], arr []int) {
	if diff := deep.Equal(list.ToArray(), arr); diff != nil {
		t.Error(diff)
	}
}

func TestSingleLinkedListInit(t *testing.T) {
	t.Run("Init", func(t *testing.T) {
		list := &SingleLinkedList[int]{}
		list.Init()

		checkListIsEmpty(t, list)
	})
}

func TestSingleLinkedListAppend(t *testing.T) {
	t.Run("Append", func(t *testing.T) {
		list := &SingleLinkedList[int]{}
		list.Init()

		checkListIsEmpty(t, list)

		list.PushBack(1)

		checkListNotEmpty(t, list, 1)
		checkListEquals(t, list, []int{1})

		list.PushBack(2)

		checkListNotEmpty(t, list, 2)
		checkListEquals(t, list, []int{1, 2})
	})
}

func TestSingleLinkedListPrepend(t *testing.T) {
	t.Run("Prepend", func(t *testing.T) {
		list := &SingleLinkedList[int]{}
		list.Init()

		checkListIsEmpty(t, list)

		list.PushFront(1)

		checkListNotEmpty(t, list, 1)
		checkListEquals(t, list, []int{1})

		list.PushFront(2)

		checkListNotEmpty(t, list, 2)
		checkListEquals(t, list, []int{2, 1})
	})
}

func TestSingleLinkedListRemove(t *testing.T) {
	t.Run("Remove", func(t *testing.T) {
		list := &SingleLinkedList[int]{}
		list.Init()

		checkListIsEmpty(t, list)

		list.PushBack(1)
		list.PushBack(2)

		checkListNotEmpty(t, list, 2)
		checkListEquals(t, list, []int{1, 2})

		list.PopFront()

		checkListNotEmpty(t, list, 1)
		checkListEquals(t, list, []int{2})

		list.PopFront()

		checkListIsEmpty(t, list)

		list.PopFront()

		checkListIsEmpty(t, list)
	})
}

func TestSingleLinkedListIterator(t *testing.T) {
	t.Run("Iterator", func(t *testing.T) {
		list := &SingleLinkedList[int]{}
		list.Init()

		checkListIsEmpty(t, list)

		list.PushBack(1)
		list.PushBack(2)

		checkListNotEmpty(t, list, 2)
		checkListEquals(t, list, []int{1, 2})

		iterator := SingleLinkedListIterator[int]{}
		iterator.Init(list)

		for i := 0; i < 2; i++ {
			if !iterator.Next() {
				t.Errorf("iterator.Next() = false; want true")
			}
			if iterator.Value() != i+1 {
				t.Errorf("iterator.Value() = %v; want %v", iterator.Value(), i+1)
			}
		}

		if iterator.Next() {
			t.Errorf("iterator.Next() = true; want false")
		}
	})
}

func BenchmarkSingleLinkedListAddDelete(b *testing.B) {
	list := &SingleLinkedList[int]{}
	list.Init()
	for i := 0; i < b.N; i++ {
		list.PushBack(i)
	}
	for i := 0; i < b.N; i++ {
		list.PopFront()
	}
}
