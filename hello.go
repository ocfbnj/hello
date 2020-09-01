package main

import (
	"container/list"
	"fmt"

	"github.com/ocfBNj/hello/listutils"
)

func main() {
	l := list.New()

	for i := 0; i != 10; i++ {
		l.PushBack(i)
	}

	listutils.RemoveOdd(l)
	listutils.PrintList(l)

	listutils.RemoveEven(l)
	listutils.PrintList(l)

	fmt.Println(l.Len())
}
