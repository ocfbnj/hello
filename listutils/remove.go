package listutils

import "container/list"

// RemoveOdd Remove odd numbers
func RemoveOdd(l *list.List) {
	remove(l, func(value interface{}) bool {
		return value.(int)&1 == 1
	})
}

// RemoveEven Remove even numbers
func RemoveEven(l *list.List) {
	remove(l, func(value interface{}) bool {
		return value.(int)&1 == 0
	})
}

func remove(l *list.List, pred func(interface{}) bool) {
	var prev *list.Element

	for cur := l.Front(); cur != nil; {
		if pred(cur.Value) {
			l.Remove(cur)

			if prev != nil {
				cur = prev
			} else {
				prev = nil
				cur = l.Front()
			}
		} else {
			prev = cur
			cur = cur.Next()
		}
	}
}
