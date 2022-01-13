package main

import "fmt"

// && = and
// || = or
// ! = negation

// fallthrough = para que el case del switch statement pase a la siguiente linea y siga leyendo

func main() {

	if nombre, edad := "Alvares", 15; edad < 14 {
		fmt.Println(nombre, "Es un niño")
	} else if edad < 18 {
		fmt.Println(nombre, "Es un adolecente")
	} else {
		fmt.Println(nombre, "Es un adulto")
	}

	var a int8 = 3
	switch a {
	case 1:
		fmt.Println("Es 1")
	case 2:
		fmt.Println("Es 2")
	default:
		fmt.Println("No se que es")
	}

	switch dia := 7; {
	case dia >= 1 && dia <= 6:
		fmt.Println("Dia de semana")
	case dia == 7 || dia == 8:
		fmt.Println("Fin de semana")
	default:
		fmt.Println("ingrese un dia valido")
	}

	fmt.Println("Fin del programa!")
}
