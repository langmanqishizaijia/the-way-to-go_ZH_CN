package main

import (
	"fmt"
	"strings"
)

func main() {
	var origS string = "Hi there! "
	var newS string

	newS = strings.Repeat(origS, 3)
	fmt.Printf("The new repeated string is: %s\n", newS)
}

func init() {
	origS := "His there! "
	newS := ""

	newS = strings.Repeat(origS, 4)
	fmt.Printf("the new repeated string is: %s\n", newS)
}
