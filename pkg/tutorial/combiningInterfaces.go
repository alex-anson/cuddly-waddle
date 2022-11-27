package tutorial

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

type WriterCloser interface {
	Writer // (this interface is declared in `interfacesExample.go`)
	Closer
}

type BufferedWriterCloser struct {
	buffer *bytes.Buffer
}

func (bwc *BufferedWriterCloser) Write(data []byte) (int, error) {
	n, err := bwc.buffer.Write(data)
	if err != nil {
		return 0, err
	}

	v := make([]byte, 8)
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

func (bwc *BufferedWriterCloser) Close() error {
	for bwc.buffer.Len() > 0 {
		data := bwc.buffer.Next(8)
		_, err := fmt.Println((string(data)))
		if err != nil {
			return err
		}
	}
	return nil
}

func NewBufferedWriterCloser() *BufferedWriterCloser {
	return &BufferedWriterCloser{
		buffer: bytes.NewBuffer([]byte{}),
	}
}
