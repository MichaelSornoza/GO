package main

import "fmt"

func main() {
	var str string = "Hola Mundo"
	a := "Gooland Mundo"

	fmt.Println(a)
	fmt.Println(str)

	num := 10.
	var b float64 = 3

	fmt.Println((num / b))

	d := 3
	var c int = 10
	fmt.Println((c / d))

	imprimirHola()
}

func privado() {

}

func Publica() {

}

func imprimirHola() {
	defer fmt.Println("Hola") //Defer indica que se debe ejecutar al final de la funcion
	fmt.Println("Mundo")
}
