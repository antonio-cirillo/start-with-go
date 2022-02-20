package main

import (
	"fmt"
)

const (
	isAdmin = 1 << iota
	isHeadquarters
	canSeeFinancials
	canSeeAfrica
	canSeeAsia
	canSeeEurope
	canSeeNorthAmerica
	canSeeSouthAmerica
)

func main() {

	fmt.Printf("isAdmin: %b\n", isAdmin)                       // 1
	fmt.Printf("isHeadquarters: %b\n", isHeadquarters)         // 10
	fmt.Printf("canSeeFinancials: %b\n", canSeeFinancials)     // 100
	fmt.Printf("canSeeAfrica: %b\n", canSeeAfrica)             // 1000
	fmt.Printf("canSeeAsia: %b\n", canSeeAsia)                 // 10000
	fmt.Printf("canSeeEurope: %b\n", canSeeEurope)             // 100000
	fmt.Printf("canSeeNorthAmerica: %b\n", canSeeNorthAmerica) // 1000000
	fmt.Printf("canSeeSouthAmerica: %b\n", canSeeSouthAmerica) // 10000000

	var roles byte = isAdmin | canSeeFinancials | canSeeEurope

	fmt.Printf("Is Admin? %v\n", roles&isAdmin == isAdmin)                      // True
	fmt.Printf("Is Headquarters? %v\n", roles&isHeadquarters == isHeadquarters) // False
}
