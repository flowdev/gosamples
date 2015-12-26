// tag::first[]
package main

type GreetAny struct {
	outPort func(string)            // <1>
}

func NewGreetAny() *GreetAny {
	return &GreetAny{}
}

func (ga *GreetAny) InPort(name string) {    // <2>
	ga.outPort("Hello, " + name + "!");
}
// end::first[]

// tag::second[]
func (ga *GreetAny) SetOutPort(outPort func(string)) {
	ga.outPort = outPort
}
// end::second[]
