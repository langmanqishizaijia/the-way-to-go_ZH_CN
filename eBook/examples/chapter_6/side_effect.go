// side_effect.go
package chapter_6

import (
	"fmt"
)

func Multiply(a, b int, reply *int) {
	*reply = a * b
}

func main() {
	n := 0
	reply := &n
	Multiply(10, 5, reply)
	fmt.Println("Multiply:", *reply) // Multiply: 50
	fmt.Println("Multiply:", n)      // Multiply: 50
}
