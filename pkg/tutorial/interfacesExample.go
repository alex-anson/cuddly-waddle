package tutorial

import "fmt"

// EXPLORING INTERFACES

func interfaceMain() {
	fmt.Println("----------------")

	// "w" is of type "Writer"... setting that equal to a ConsoleWriter instance.
	// So "w" is holding a Writer, which is something that implements the Writer
	// interface. We don't know the concrete type - we know how to invoke the Write
	// method because that's defined by the Writer interface, but within this
	// interfaceMain ƒn, we don't know what is being written to - that's the
	// responsibility of the actual implementation (the Write method)
	var w Writer = ConsoleWriter{}
	w.Write([]byte("hey golang"))

	// My own learning purposes
	var w2 Writer = AnotherConsoleWriter{}
	w2.Write([]byte("cool. 😎"))

	myInt := IntCounter(0)
	var inc Incrementer = &myInt

	for i := 0; i < 6; i++ {
		fmt.Println(inc.Increment())
	}
}

// If this were a struct, we'd enter the data we want the struct to hold on to -
// structs are ultimately data containers. Interfaces don't describe data, interfaces
// describe behavior. (... method definitions)
// DEFINE THE INTERFACE
// Anything that implements this interface will take in a slice of bytes and write
// it to something (i.e. the console, a TCP connection, the file system).
type Writer interface {
	// Will return an integer (normally, the number of bytes written) and an error
	Write([]byte) (int, error)
}

// Gonna implement the Writer interface with a ConsoleWriter implementation, which
// is a struct
type ConsoleWriter struct{}

// In Go, we don't explicitly implement interfaces.

// Implicitly implement the interface by creating a method on the ConsoleWriter
// that has the signature of a Writer interface.
func (cw ConsoleWriter) Write(data []byte) (int, error) {
	// Without wrapping 'data' in 'string()', you'll get the byte slice printed
	n, err := fmt.Println(string(data))
	return n, err
}

// For my own learning purposes
type AnotherConsoleWriter struct{}

func (cw AnotherConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println("(pretend this is the file system) ... ", string(data))
	return n, err
}

// Don't HAVE to use a struct (though it's one of the most common ways to implement
// interfaces). Any type that can have a method associated with it (this is ALL
// types) can implement an interface.
// Defining a new interface, Incrementer:
type Incrementer interface {
	// Increment is a method that returns an integer
	Increment() int
}

type IntCounter int

func (ic *IntCounter) Increment() int {
	*ic = *ic + 2
	return int(*ic)
}
