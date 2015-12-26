package  main

import (
    "fmt"
    "os"
)

type OutputText struct {
}

func NewOutputText() *OutputText {
	return &OutputText{}
}

func (ot *OutputText) InPort(text string) {    // <3>
	fmt.Fprintln(os.Stderr, text)
}
