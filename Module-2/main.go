package main

import "fmt"

func main() {
	//!Arreglos
	/*var numeros [5]int
	numeros[0] = 100
	fmt.Println(numeros)*/

	//!Arreglo inicializado (los tres puntos indican que el sistema debe determinar automaticamente cual es el tamaño del arreglo)
	/*numeros := [...]int{100, 200, 300, 400, 500}
	fmt.Println(numeros)*/

	//!Asignar un indice al arreglo
	/*monedas := [...]string{0: "Peso Colombiano", 2: "Dolar", 1: "Euro"}
	fmt.Println(monedas)*/

	//!Slices
	/*numeros := []int{1, 2, 3, 4, 5}
	numeros = append(numeros, 6)
	numeros = append(numeros, 7)
	numeros = append(numeros, 8)

	nuevoSlice := numeros[0:5]
	//* Si se modifica el Slice numero, se modifica el Slice nuevoSlice
	//* Los Slice son referencias de arreglos previamente establecidos
	fmt.Println(numeros)
	fmt.Println(nuevoSlice)*/

	//!Caracteristicas slices
	/*meses := []string{"Enero", "Febrero", "Marzo", "Abril"}
	longitud := len(meses)
	capacidad := cap(meses)

	fmt.Println(longitud)
	fmt.Println(capacidad)

	//*Si la estructura se encuentra en su capacidad maxima, se genera un nuevo arreglo
	//* Puntero: Donde comienza la subsecuencia
	//* Longitud: Cantidad de elementos dentro de un slice
	//* Capacidad: Cantidad maxima de elementos que almacena un slice

	meses = append(meses, "Mayo")
	longitud = len(meses)
	capacidad = cap(meses)

	fmt.Println(longitud)
	fmt.Println(capacidad)*/

	//!Funcion make
	//* make (arreglo, longitud, capacidad)
	/*slice := make([]int, 3, 3)
	slice[0] = 100
	slice[1] = 200
	slice[2] = 300

	slice = append(slice, 400)

	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))*/

	//!Mapas
	/*dias := make(map[int]string)
	dias[0] = "Domingo"
	dias[1] = "Lunes"
	dias[2] = "Martes"
	dias[3] = "Miercoles"

	dias[1] = "LUNES"
	delete(dias, 3)
	fmt.Println(dias)*/

	//* Es posible crear estructuras mas complejas
	/*usuarios := make(map[string][]int)
	usuarios["Eduardo"] = []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	fmt.Println(usuarios)*/

	//*¿Como iterar un mapa?
	usuarios := map[int]string{}
	usuarios[1] = "Usuario 1"
	usuarios[2] = "Usuario 2"
	usuarios[3] = "Usuario 3"
	usuarios[4] = "Usuario 4"

	for llave, valor := range usuarios {
		fmt.Println(llave, valor)
	}
}
