package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	inicio := time.Now()

	servidores := []string{
		"http://platzi.com",
		"http://google.com",
		"http://186.42.119.139",
	}
	for _, servidor := range servidores {
		revisarServidor(servidor)
	}

	tiempoPaso := time.Since(inicio)
	fmt.Printf("Tiempo de Ejecucion %s\n", tiempoPaso)
}

func revisarServidor(servidor string) {
	_, err := http.Get(servidor)

	if err != nil {
		fmt.Println(servidor, " no esta disponible")
	} else {
		fmt.Println(servidor, " Operativo")
	}
}
