package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {
	var w Writer = ConsoleWriter{}
	w.Write([]byte("hello Go"))
	myInt := IntCounter(0)
	var inc Incrementer = &myInt
	for i := 0; i < 10; i++ {
		fmt.Println(inc.Increment())
	}
	var wc WriterCloser = NewBufferWriterCloser()
	wc.Write([]byte("hello Youtube listeners, this is test"))
	defer wc.Close()
	r, ok := wc.(*BufferedWriterCloser)
	if ok {
		fmt.Println(r)
	} else {
		fmt.Println("conversion failed")
	}
	var myObj interface{} = NewBufferWriterCloser()
	if wc1, ok := myObj.(WriterCloser); ok {
		wc1.Write([]byte("Hello Youtube listeners, this is a test"))
		wc1.Close()
	}
	r1, ok1 := myObj.(io.Reader)
	if ok1 {
		fmt.Println(r1)
	} else {
		fmt.Println("Conversion failed")
	}

	var i interface{} = true
	switch i.(type) {
	case int:
		fmt.Println("i is an integer")
	case string:
		fmt.Println("i is a string")
	case bool:
		fmt.Println("i is a bool")
	default:
		fmt.Println("I don't know what i is")
	}
}

type Writer interface {
	Write([]byte) (int, error)
}

type Closer interface {
	Close() error
}

type WriterCloser interface {
	Writer
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
		_, err := fmt.Println(string(data))
		if err != nil {
			return err
		}
	}
	return nil
}
func NewBufferWriterCloser() *BufferedWriterCloser {
	return &BufferedWriterCloser{
		buffer: bytes.NewBuffer([]byte{}),
	}
}

type ConsoleWriter struct{}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}

type Incrementer interface {
	Increment() int
}
type IntCounter int

func (ic *IntCounter) Increment() int {
	*ic++
	return int(*ic)
}
