package main

import (
	"fmt"
)

type greeter struct {
	greeting string
	name     string
}

func (g greeter) greet() {
	fmt.Println(g.greeting, g.name)
}

func (g *greeter) setName(name string) {
	g.name = name
}

func main() {

	g := greeter{
		greeting: "Hello",
		name:     "Go",
	}

	g.greet()
	fmt.Println("Name:", g.name)
	g.setName("GoLang")
	fmt.Println("Name:", g.name)

}
