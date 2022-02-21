package main

import (
	"fmt"
)

type Animal struct {
	Name   string
	Origin string
}

type Bird struct {
	Animal   // Composition
	SpeedKPH float32
	CanFly   bool
}

func main() {

	b := Bird{
		Animal{"Emu", "Australia"},
		48,
		false,
	}

	/*
	* b.Name = "Emu"
	* b.Origin = "Australia"
	* b.SpeedKPH = 48
	* b.CanFly = false
	 */

	fmt.Println(b)

}
