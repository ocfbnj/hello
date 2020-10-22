package main

import (
	"fmt"
)

func main() {
	str := "你好"

	for _, v := range str {
		fmt.Printf("0x%x\n", v)
	}
}
