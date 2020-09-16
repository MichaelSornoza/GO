package main

import "fmt"

func main() {
	x := 25
	fmt.Println(x)
	//IMpreme donde esta en memoria
	fmt.Println(&x)
	// cambiarValor(x)

	y := &x
	//Imprime donde esta en memoria
	fmt.Println(y)
	//Imprime el valor almacenado en ese espacio de memoria
	fmt.Println(*y)

}

func cambiarValor(a int) {
	fmt.Println(&a)
	a = 36
}
