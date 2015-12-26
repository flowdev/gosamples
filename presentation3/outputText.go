package  main

import (
    "fmt"
    "os"
)

type OutputText struct {
	getText func(string)string     // <1>
}

func NewOutputText(getText func(string)string) *OutputText {     // <2>
	return &OutputText{ getText }
}

func (ot *OutputText) InPort(data string) {
	fmt.Fprintln(os.Stderr, ot.getText(data))    // <3>
}
