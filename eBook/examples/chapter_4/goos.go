package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	var goos string = runtime.GOOS
	fmt.Printf("The operating system is: %s\n", goos)
	path := os.Getenv("PATH")
	fmt.Printf("Path is %s\n", path)
}

func init() {
	var goos string = runtime.GOOS
	fmt.Printf("The operating system is: %s\n", goos)

	path := os.Getenv("PATH")
	fmt.Printf("Path is %v\n", path)

	fmt.Printf("gid is %v\n", os.Getegid())
}
