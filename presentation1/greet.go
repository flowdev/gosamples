// tag::first[]
package main

type Greet struct {
	outPort func(string)            // <1>
}

func NewGreet() *Greet {
	return &Greet{}
}

func (g *Greet) InPort(data []string) {    // <2>
	for i := 0; i < len(data); i++ {
		g.outPort("Hello, " + data[i] + "!");
	}
}
// end::first[]

// tag::second[]
func (g *Greet) SetOutPort(outPort func(string)) {
	g.outPort = outPort
}
// end::second[]
