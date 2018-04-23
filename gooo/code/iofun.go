package main

import(
	"io"
	"os"
	"fmt"
)

type Foo interface {
	Read([]byte) (int, error)
	Close() error
}

func useReadCloser(rc io.ReadCloser) {
	fmt.Println("useReadCloser")
}

func useReader(r io.Reader) {
	fmt.Println("useReader")
}

func useFoo(f Foo) {
	fmt.Println("useFoo")
}

func main() {
}