package main

import "fmt"

func main() {
	//bool, string, numeric

	var t int;

	var x uint8 = 200;
	var y uint16 = 2000;

	//casting y ademas _ para cuando no vas a usar aun 
	//una variable
	_ = uint16(x) + y; 

	fmt.Printf("Tipo: %T, Valor: %v\n", t, t);
}
