// test_defer.go
package chapter_6

import (
	"fmt"
)

func f() (ret int) {
	defer func() {
		ret++
	}()
	return 1
}

func main() {
	fmt.Println(f())
}

// Output: 2
