package main

import "fmt"

type taskList struct {
	//Arreglos con tamanos fijos / No se puede hacer CRUD
	//Slices son Listas - El siguiente es un slice
	tasks []*task
}

func (t *taskList) agregarALista(tarea *task) {
	//Append Toma un Slice para agregar al slice principal
	t.tasks = append(t.tasks, tarea)
}

type task struct {
	nombre      string
	descripcion string
	completado  bool
}

func (t task) marcarCompleta() {

}

func (t *task) completo() {
	t.completado = true
}

func (t *task) actualizarDescripcion(descripcion string) {
	t.descripcion = descripcion
}

func (t *task) actualizarNombre(nombre string) {
	t.nombre = nombre
}

func (t *taskList) borrarTask(index int) {
	// Eliminando un elemento del slice
	// [:index] Desde el inicio hasta el índice
	// [index+1:] Desde el índice hasta el final
	t.tasks = append(t.tasks[:index], t.tasks[index+1:]...)
}

func main() {
	//Enviamos referencia de el struct y manda el espacio en memoria
	t1 := &task{
		nombre:      "Completar mi curso de Go",
		descripcion: "Completar mi curso de Go de Platzi en esta semana",
	}
	t2 := &task{
		nombre:      "Completar mi curso de Python",
		descripcion: "Completar mi curso de Python de Platzi en esta semana",
	}
	t3 := &task{
		nombre:      "Completar mi curso de JavaScript",
		descripcion: "Completar mi curso de JavaScript de Platzi en esta semana",
	}

	lista := &taskList{
		//Declarar Slice
		tasks: []*task{
			t1, t2,
		},
	}
	lista.agregarALista(t3)

	// for i := 0; i < len(lista.tasks); i++ {
	// 	fmt.Println("Index", i, "nombre", lista.tasks[i].nombre)
	// }

	// //Forma Elegante de For
	// //Range me devuelve el indice y el valor
	// //Forma mas utilizada, es como un map en JavaScript
	// for index, tarea := range lista.tasks {
	// 	fmt.Println("Index", index, "nombre", tarea.nombre)
	// }

	for i := 0; i < 10; i++ {

		if i == 5 {
			//Cancela la iteracion
			break
		}
		fmt.Println(i)
	}
	for i := 0; i < 10; i++ {

		if i == 5 {
			//Cancela la iteracion
			continue
		}
		fmt.Println(i)
	}
}

// 	fmt.Println(len(lista.tasks))

// 	fmt.Println(lista.tasks[0])

// 	//Len es para ver el tamano del Slice
// 	fmt.Println(len(lista.tasks))

// 	lista.borrarTask(1)
// 	fmt.Println(len(lista.tasks))
// 	//Formato para imprimir cual es mi propiedad y el valor e la misma
// 	// fmt.Printf("%+v\n", t)
// 	// t.marcarCompleta()
// 	// t.actualizarNombre("Finalizar mi curso de Go")
// 	// t.actualizarDescripcion("Completar mi curso cuanto antes")

// 	//fmt.Printf("%+v\n", t)
// }
