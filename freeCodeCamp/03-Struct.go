package main

import "fmt"

type Doctor struct {
	number     int
	actorName  string
	companions []string
}

func main() {

	aDoctor := Doctor{
		number:    3,
		actorName: "John Pertwee",
		companions: []string{
			"Liz Shaw",
			"Jo Grant",
			"Sarah Jane Smith",
		},
	}

	fmt.Println(aDoctor)
	fmt.Println("Actor Name: " + aDoctor.actorName)

	anotherDoctor := struct{ name string }{name: "John Pertwee"}
	fmt.Println(anotherDoctor)

}
