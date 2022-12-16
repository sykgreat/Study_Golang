package _stack_pointer

import (
	"testing"
)

func Test_Main1(t *testing.T) {

	// Declare variable of type int with a value of 10.
	count := 10

	// Display the "value of" and "address of" count.
	println("count:\tValue Of[", count, "]\tAddr Of[", &count, "]")

	// Pass the "value of" the count.
	increment1(count)

	println("count:\tValue Of[", count, "]\tAddr Of[", &count, "]")
}

//go:noinline
func increment1(inc int) {

	// Increment the "value of" inc.
	inc++
	println("inc:\tValue Of[", inc, "]\tAddr Of[", &inc, "]")
}

func Test_Main2(t *testing.T) {

	// Declare variable of type int with a value of 10.
	count := 10

	// Display the "value of" and "address of" count.
	println("count:\tValue Of[", count, "]\t\tAddr Of[", &count, "]")

	// Pass the "address of" count.
	increment2(&count)

	println("count:\tValue Of[", count, "]\t\tAddr Of[", &count, "]")
}

//go:noinline
func increment2(inc *int) {

	// Increment the "value of" count that the "pointer points to". (dereferencing)
	*inc++
	println("inc:\tValue Of[", inc, "]\tAddr Of[", &inc, "]\tValue Points To[", *inc, "]")
}
