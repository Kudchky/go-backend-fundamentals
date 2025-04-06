package main 

import "fmt"

func main()  {
    if character := "scarface"; character == "batman" {
        fmt.Print("Genial, el personaje es un superheroe!\n")
    } else if character == "scarface" {
        fmt.Print("Oh no!, tu personaje se un villano\n")
    } else {
        fmt.Print("Es un personaje normal\n")
    }

}
