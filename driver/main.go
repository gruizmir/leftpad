package main

import (
	"fmt"

	"github.com/gruizmir/leftpad"
)

func main() {
	fmt.Println(leftpad.Format("hello", 15))
	fmt.Println(leftpad.FormatRune("hello", 15, '-'))
}
