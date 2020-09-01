package listutils

import (
	"container/list"
	"fmt"
)

// PrintList Print elements of the l
func PrintList(l *list.List) {
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Printf("%v ", e.Value)
	}

	fmt.Println()
}
