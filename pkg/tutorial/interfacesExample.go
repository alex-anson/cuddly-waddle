package tutorial

import (
	"fmt"
	"io"
)

// EXPLORING INTERFACES
// Implicit implementation.

func interfaceMain() {
	fmt.Println("----------------")

	// "w" is of type "Writer"... setting that equal to a ConsoleWriter instance.
	// So "w" is holding a Writer, which is something that implements the Writer
	// interface. We don't know the concrete type - we know how to invoke the Write
	// method because that's defined by the Writer interface, but within this
	// interfaceMain ƒn, we don't know what is being written to - that's the
	// responsibility of the actual implementation (the Write method that's associated
	// with the ConsoleWriter)
	var w Writer = ConsoleWriter{}
	w.Write([]byte("hey golang"))

	// Polymorphic behavior - w.Write(...) doesn't care what it's writing to.

	// My own learning porpoises 🐬
	var w2 Writer = AnotherConsoleWriter{}
	w2.Write([]byte("cool. 😎"))

	// Create integer counter - cast an integer (0) to an IntCounter in order to do that.
	myInt := IntCounter(0)
	// Create the incrementer, assigned to a pointer of the myInt object.
	var inc Incrementer = &myInt

	for i := 0; i < 6; i++ {
		fmt.Println(inc.Increment())
	}

	// Part of the example from `combiningInterfaces.go`
	var wc WriterCloser = NewBufferedWriterCloser()
	// Converted string to byte slice because that's what the Write method expects.
	wc.Write([]byte("uh, this is a test. let's make it longer, just so what's happening is more obvious"))
	wc.Close()

	// Type conversion.
	// storeIt := object.(typeToConvertTo)     ... asterisk needed if it's a pointer type
	bwc := wc.(*BufferedWriterCloser)
	fmt.Println(bwc) // Memory address of this BufferedWriterCloser
	// If the type conversion succeeds, you can work with that variable now. In this
	// case, could work with bwc no longer as a WriterCloser, but as a BufferedWriterCloser.
	// ......
	// If the type conversion fails, the app would panic. DON'T make panic & recovery
	// primary control flows in an app. panic is expensive
	// instead, could:
	r, ok := wc.(io.Reader)
	if ok {
		fmt.Println(r)
	} else {
		fmt.Println("conversion failed")
	}

	// Continuing type conversion...
	// Empty interface is just that, an interface with no methods on it. What's
	// nice about it is that *everything* can be cast into an object that has no
	// methods on it, even primitives. (i.e. can cast an integer to an emtpy interface)
	// Useful in situations where you have multiple things you need to work with
	// that aren't type compatible
	// Problem with empty interface: can't do anything with it. Because it has no
	// exposed methods.
	// In order to do something useful with a variable that has the type of an empty
	// interface, you'll need to do a type conversion, or you'll need to start using
	// the "reflect" package in order to figure out what kind of object you're dealing with.
	var someObject interface{} = NewBufferedWriterCloser()
	// ^ Defining an interface on the fly
	// Trying to type cast into a WriterCloser
	if wc, ok := someObject.(WriterCloser); ok {
		wc.Write([]byte("still works"))
		wc.Close()
	}
	// Empty interface is very common. BUT - it's almost always gonna be an intermediate step.

	// Type switches
	var i interface{} = 0
	switch i.(type) {
	// ^ this is a type switch.
	// the cases are types.
	// Type switches are commonly paired with the empty interface.
	case int:
		fmt.Println("logic to process an integer goes here")
	case string:
		fmt.Println("logic to process a string goes here")
	default:
		fmt.Println("logic to process something else goes here")
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

// ^ The naming convention for single method interfaces (which are common) is to
// use the method name plus "er". (i.e., a Read method's interface should be Reader)
// Naming multiple method interfaces is more challenging, but it should always be
// named by what it does.

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

// Creating a method on the AnotherConsoleWriter struct that has the signature of
// a Writer instance.
func (cw AnotherConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println("(pretend this is the file system) ... ", string(data))
	return n, err
}

// Don't HAVE to use a struct (though it's one of the most common ways to implement
// interfaces). Any type that can have a method associated with it (this is ALL
// types) can implement an interface.
// Defining a new interface, Incrementer:
type Incrementer interface {
	// Increment will be a method that returns an integer. It will increment
	// *something.* So, whatever we're gonna implement this thing with, is going
	// to increment values.
	Increment() int
}

// Defining a type alias for an integer.
type IntCounter int

// ^ Can't add a method directly to the int type - that is not under our control
// i.e., couldn't do "func (ic *int) Increment() int {" in the implementation --
// error = "cannot define new methods on non-local type int        compiler InvalidRecv"

// Adding a method to the IntCounter custom type. This is the implementation for
// the Incrementer interface.
func (ic *IntCounter) Increment() int {
	// Incrementing the type itself... "type IntCounter int" is a type alias for
	// an integer - which IS A NUMBER.
	// We have a type defined on an integer, and the integer itself is storing the
	// data that the method is using.
	*ic = *ic + 2
	return int(*ic)
}
