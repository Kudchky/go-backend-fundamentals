package main

import (
	"fmt"
	"math"
)

func areaCirculo(radio float64) float64 {
	return math.Pi * radio * radio
}

func main() {

	var radio float64


	fmt.Printf("Por favor, ingrese el radio del circulo: ")
	if _, err := fmt.Scanln(&radio); err != nil {
		fmt.Println("Error, por favor ingrese un valor valido para radio", err)
		return
	}

	if radio <= 0 {
		fmt.Println("Error, el radio debe ser positivo")
		return
	}

	area := areaCirculo(radio)

	fmt.Printf("El area del circulo es: %.2f\n", area)
}
