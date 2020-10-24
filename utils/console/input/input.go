package input

import (
	"fmt"
)

// Read an Integer value from the console
// The "template" is the "prefix" text to help signalize to the user we are
// Waiting on the input
func ReadInt(template string) int {
	var ret int
	print(template)
	fmt.Scanf("%d", &ret)
	return ret
}