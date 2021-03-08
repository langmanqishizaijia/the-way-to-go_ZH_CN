// function_parameter.go
package main

import (
	"fmt"
)

func main() {
	callback(1, 2, Add)
	callback(2, 1, Sub)
}

func Sub(a, b int) {
	fmt.Printf("The sub of a=%v, b=%v, sub=%v", a, b, a-b)
}
func Add(a, b int) {
	fmt.Printf("The sum of %d and %d is: %d\n", a, b, a+b)
}

func callback(x, y int, f func(int, int)) {
	f(x, y) // this becomes Add(1, 2)
}

// Output:  The sum of 1 and 2 is: 3
