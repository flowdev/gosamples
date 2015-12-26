// tag::first[]
package main

import "os"

type SayHello struct {
	greet *Greet            // <1>
	outputText *OutputText
	InPort func([]string)   // <2>
}
// end::first[]

// tag::second[]
func NewSayHello() *SayHello {
	sh := &SayHello{NewGreet(), NewOutputText(), nil}      // <1>

	sh.greet.SetOutPort(sh.outputText.InPort)    // <2>
	sh.InPort = sh.greet.InPort        // <3>

	return sh
}
// end::second[]

// tag::third[]
func main() {
	sh := NewSayHello()
	sh.InPort(os.Args[1:])
}
// end::third[]
