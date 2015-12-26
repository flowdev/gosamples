// tag::first[]
package main

type GiveNames struct {
	outPort func(string)            // <1>
}

func NewGiveNames() *GiveNames {
	return &GiveNames{}
}

func (g *GiveNames) InPort(names []string) {    // <2>
	for i := 0; i < len(names); i++ {
		g.outPort(names[i]);
	}
}
// end::first[]

// tag::second[]
func (g *GiveNames) SetOutPort(outPort func(string)) {
	g.outPort = outPort
}
// end::second[]
