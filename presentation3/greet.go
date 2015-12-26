// tag::first[]
package main

type Greet struct {
	giveNames *GiveNames
	greetAny *GreetAny
	InPort func([]string)
}
// end::first[]

// tag::second[]
func NewGreet() *Greet {
	g := &Greet{NewGiveNames(), NewGreetAny(), nil}

	g.giveNames.SetOutPort(g.greetAny.InPort)
	g.InPort = g.giveNames.InPort

	return g
}

func (g *Greet) SetOutPort(outPort func(string)) {
	g.greetAny.SetOutPort(outPort)
}
// end::second[]
