package main

import (
	"fmt"
	"github.com/superdbop/learning-go/stringutil"
)

func main() {
	fmt.Printf("hello, world!\n")
	fmt.Printf(stringutil.Reverse("hello, world!") + "\n")
}
