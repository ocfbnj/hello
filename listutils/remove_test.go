package listutils

import (
	"container/list"
	"testing"
)

func TestRemoveOdd(t *testing.T) {
	l := list.New()

	for i := 0; i != 10; i++ {
		l.PushBack(i)
	}

	RemoveOdd(l)

	for e := l.Front(); e != nil; e = e.Next() {
		if e.Value.(int)&1 == 1 {
			t.Errorf("Can not remove all odd numbers: %v", e.Value.(int))
		}
	}
}

func TestRemoveEven(t *testing.T) {
	l := list.New()

	for i := 0; i != 10; i++ {
		l.PushBack(i)
	}

	RemoveEven(l)

	for e := l.Front(); e != nil; e = e.Next() {
		if e.Value.(int)&1 == 0 {
			t.Errorf("Can not remove all even numbers: %v", e.Value.(int))
		}
	}
}
