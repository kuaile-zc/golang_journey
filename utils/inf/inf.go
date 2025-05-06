package inf

import (
	"fmt"
	"io"
	"os"
)

type ByteCounter int

func (bc *ByteCounter) Write(p []byte) (int, error) {
	*bc += ByteCounter(len(p))
	return len(p), nil
}

func InterfaceTest() {
	var bc ByteCounter
	bc.Write([]byte("Hello"))
	fmt.Println(bc)
	bc = 0
	var name = "Dolly"
	fmt.Fprintf(&bc, "Hello, %s", name)
	fmt.Println(bc)

	fmt.Println("-------")
	var w io.Writer
	w = os.Stdout
	w.Write([]byte("Hello"))

	fmt.Println("\n-------")

}
