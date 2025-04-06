package main

import "fmt"

// const os, domain = "linux", "EDteam.com"
// Las constantes se deben trabajar a nivel de paquete
// o como en otros lenguajes a nivel global

// No olvidar el operador iota para valores que son seguidos

const (
	os     = "linux"
	domain = "EDTeam.com"
)

const (
	Jan = iota + 1
	Feb
	Mar
	Abr
	May
)
func main() {

	fmt.Println(os, domain, May)
}
