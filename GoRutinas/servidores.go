package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	inicio := time.Now()

	canal := make(chan string)

	servidores := []string{
		"http://platzi.com",
		"http://google.com",
		"http://instagram.com",
		"http://facebook.com",
		"http://seguridadciudadana.manta.gob.ec",
		"http://portalciudadano.manta.gob.ec",
		"http://sistemasic.manta.gob.ec",
		"http://186.42.119.139",
	}

	i := 0

	for {
		if i > 10 {
			break
		}

		for _, servidor := range servidores {
			go revisarServidor(servidor, canal)
		}

		time.Sleep(1 * time.Second)
		fmt.Println(<-canal)

		i++
	}

	/*for i := 0; i < len(servidores); i++ {
		fmt.Println(<-canal)
	}*/

	tiempoPaso := time.Since(inicio)
	fmt.Printf("Tiempo de Ejecucion %s\n", tiempoPaso)
}

func revisarServidor(servidor string, canal chan string) {
	_, err := http.Get(servidor)

	if err != nil {
		canal <- servidor + " no se encuentra disponible"
	} else {
		canal <- servidor + " se encuentra operativo"
	}
}
