package main

// pointer are used to store the memory address of another variable
// & is a address operator
// * is a dereference operator
// * is also used to create a pointer variable

func main() {
	// memory -> address -> value
	a := 10
	println(&a) //show memory address

	// pointer is a variable that stores the memory address of another variable
	var pointer *int = &a // * is a address operator
	*pointer = 20         // * is a dereference operator
	println(a)

	b := &a // b is address of a (pointer)
	*b = 30
	println(*b) // *b is value of a
	println(a)
}
