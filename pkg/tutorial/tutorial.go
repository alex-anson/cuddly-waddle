package tutorial

import (
	"fmt"
	"strconv"
)

// Working with methods.

func Tutorial() {
	// Declaring a greeter struct
	g := greeter{
		greeting: "Hey",
		name:     "alex",
	}
	// Method invocation: Calling the greet ƒn, preceeding it with the above struct.
	g.greet()

	// myNumber type has a method called printInteger.
	myNumber.printInteger(12)

	// Declaring an anotherExample struct
	m := anotherExample{
		x: "Second favorite number is",
		y: 11,
	}

	// Invoke the method that is declared on the anotherExample struct.
	// When myMethod is called, myMethod gets a copy of the anotherExample object,
	// which is named "a" within the context of myMethod. Can then access the fields
	// on that object, which are "x" and "y"
	m.myMethod()
}

type greeter struct {
	greeting string
	name     string
}

// Method on the greeter struct.
// "(g greeter)" is what makes this function a method.
// "A function that's executing in a known context. In Go, a known context is any
// type." ... we're using the "greeter" struct, but could use any type. It's common
// to use structs.
func (g greeter) greet() {
	fmt.Println(g.greeting, g.name)
	// Operating on a COPY of the greeter object - not using a pointer. Changes
	// here will not be reflected in the Tutorial function
}

type myNumber int

// Added this + related stuff separate from the tutorial. I got curious.
func (i myNumber) printInteger() {
	// cannot use i (variable of type myNumber) as int value in argument to
	// strconv.Itoa     compiler (IncompatibleAssign)
	// ^ To get rid of this error, needed to wrap i in int()
	fmt.Println("Favorite number is " + strconv.Itoa(int(i)))
}

type anotherExample struct {
	x string
	y int
}

// Added this + related stuff separate from the tutorial. I got curious.
// "(a anotherExample)" gives us access to the types on that struct within myMethod
// "(a anotherExample)" provides a context that "myMethod" is executing in.
// "anotherExample" is specified as a value type; not using a pointer. It's called
// a value receiver. The received object in "myMethod" is the value "anotherExample".
func (a anotherExample) myMethod() {
	fmt.Println(a.x, a.y)
}
