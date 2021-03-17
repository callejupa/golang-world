package main

import (
	"fmt"
)

var estado bool

func main() {
	estado = true
	if estado {
		fmt.Println("Entra")
	} else {
		fmt.Println("Entra al else")
	}

	//asinacion de valor dentro de la condición
	if estado = false; estado {
		fmt.Println("Entra")
	} else {
		fmt.Println("Entra al else")
	}

	//Switch
	switch numero := 4; numero {
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	case 3:
		fmt.Println(3)
	default:
		fmt.Println("Mayor a 3")
	}
}
