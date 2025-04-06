package main

import "fmt"

func main() {
	character := "wasoni"
	canSearch := true

	switch {
	case !canSearch:
		fmt.Println("no esta habilitada la busqueda")
	case character == "superman" || character == "batman":
		fmt.Println("es un superheroe")
	case character == "scarface" || character == "wason":
		fmt.Println("es un villano")
	default:
		fmt.Println("es un personaje normal")
	}
}
