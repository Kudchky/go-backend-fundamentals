package main

import "fmt"

func main() {
	var name, country string 
    var lastname string = "Juan"

    fmt.Printf("holasssss %s", lastname)

	fmt.Print("Por favor, ingrese su nombre: ")
	if _, err := fmt.Scanln(&name); err != nil	{
		fmt.Println("Error al leer el nombre: ", err)
		return
	}

	fmt.Print("Por favor, ingrese su país: ")
	if _, err := fmt.Scanln(&country); err != nil {
		fmt.Println("Error al leer el país: ", err)
		return
	}

	fmt.Printf("Hola %s, tu eres de %s. ¡Bienvenido!\n" , name, country)
    fmt.Printf("hola mano cuando")
}
