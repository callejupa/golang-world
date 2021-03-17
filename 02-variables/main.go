package main

import "fmt"

//Declaracion de variables con tipado. Primera forma
// variables declaradas globalmente en el paquete
var numero int // int8, int16, int32, int64, uint8.. (sin signos)
var texto string = "Hola, "
var status bool

//variables publicas con la primera letra en minuscula
//variables privadas solo para el paquete con la primera letra en minuscula
//funciona igual para declaracion de funciones

func main() {
	// var num 3, num4 num5 int
	// num2, num3, num4, texto := 1, 2, 3, "esto es texto" -- asignacion por orden
	// conversion numero2 = int(num3)
	// Hay una libreria llamada strconv para hacer conversiones
	numero2 := 4 //declaracion donde el tipo es definido implicitamente
	fmt.Println(numero)
	fmt.Println(texto)
	fmt.Println(status)
	fmt.Println(numero2)

	// Cuando declaro una variable sin inicializar, por defecto pone cero para int, string vacio para string y false para bool
}
