package tutorial

// The Write method for the WriterCloser interface will print a string to the console
// in chunks (lack of better term) consisting of no more than eight characters.
// NOTE: This example isn't showing an efficient way to accomplish this task, it's
// just an example of how you might use this WriterCloser interface in a way that
// runs in the Go playground easily

import (
	"bytes"
	"fmt"
)

// How to compose interfaces together - very powerful concept in Go, one of the
// keys to scalability.
// Single method interfaces are very common & powerful, BECAUSE they define a very
// specific behavior. Because they have one method, they're pretty unopinionated,
// and can be implemented in a lot of ways - which means they're really flexible.

/** The relevant portion from the main ƒn of the interfacesExample.go file --

func interfaceMain() {
	...

	var wc WriterCloser = NewBufferedWriterCloser()
	wc.Write([]byte("uh, this is a test"))
	wc.Close()
}
*/

type Closer interface {
	Close() error
}

// Interface that's composed of other interfaces
type WriterCloser interface {
	Writer // (this interface is declared in `interfacesExample.go`)
	Closer
}

// ^ WriterCloser interface is gonna be implemented if an object has the
// "Write([]byte) (int, error)" method AND the "Close() error" method on it.

type BufferedWriterCloser struct {
	buffer *bytes.Buffer
}

func (bwc *BufferedWriterCloser) Write(data []byte) (int, error) {
	// "data" is stored in the internal buffer that the BufferedWriterCloser struct
	// defines -- bwc.buffer.Write(data)
	n, err := bwc.buffer.Write(data)
	if err != nil {
		return 0, err
	}

	// Initializes a slice of bytes with a maximum length of 8
	v := make([]byte, 8)
	// What the loop is doing: if the data stored in bwc.buffer is more than 7 bytes
	// long, it'll be printed. Otherwise, it won't print anything.
	for bwc.buffer.Len() > 8 {
		_, err := bwc.buffer.Read(v)
		if err != nil {
			return 0, err
		}
		_, err = fmt.Println(string(v))
		if err != nil {
			return 0, err
		}
	}
	return n, nil
}

// "Flush the rest of the buffer" (if it isn't already empty - in cases where the
// input string was a multiple of 8)
func (bwc *BufferedWriterCloser) Close() error {
	for bwc.buffer.Len() > 0 {
		// Pulls the remaining characters out
		data := bwc.buffer.Next(8)
		// Prints them
		_, err := fmt.Println((string(data)))
		if err != nil {
			return err
		}
	}
	return nil
}

// This is a constructor function that's returning a pointer to a BufferedWriterCloser.
// Need to do this because we need to initialize the internal buffer to a "NewBuffer"
// ... "internal buffer" = the buffer property that's on the BufferedWriterCloser
// struct?? ... i think so..
func NewBufferedWriterCloser() *BufferedWriterCloser {
	return &BufferedWriterCloser{
		buffer: bytes.NewBuffer([]byte{}),
	}
}
