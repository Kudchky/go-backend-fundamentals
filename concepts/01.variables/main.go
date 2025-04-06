package main

import (
	"fmt"
	"math"
)

func calcularAreaCirculo(radio float64) float64 {
	return math.Pi * radio * radio
}

func main () {
	var radio float64

	fmt.Print("Ingrese el radio del circulo: ")
	if _, err := fmt.Scanln(&radio); err != nil {
		fmt.Println("Error: debes ingresar un numero valido (ej: 5.5)")
		return
	}

	if radio <= 0 {
		fmt.Println("Error: el radio debe ser positivo")
		return
	}

	area := calcularAreaCirculo(radio)

	fmt.Printf("El area del circulo es: %.2f\n", area)
}
