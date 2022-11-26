package tutorial

import "fmt"

func Tutorial() {
	g := greeter{
		greeting: "Hey",
		name:     "alex",
	}
	g.greet()
}

type greeter struct {
	greeting string
	name     string
}

func (g greeter) greet() {
	fmt.Println(g.greeting, g.name)
}
