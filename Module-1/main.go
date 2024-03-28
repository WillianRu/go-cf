package main

import "fmt"

// !Uso de constantes
//const Curso string = "Go para profesionales"

//!Secuencia de valores
/*const (
	Domingo int = iota + 1 //1
	Lunes
	Martes
	Miercoles
)*/

func main() {

	//!Variables
	/*var nombre string
	nombre = "Willian Rueda"

	var edad int = 26*/
	//!Autoasignacion de valor y tipo de variable
	/*nombre := "Willian"
	edad := 25
	altura := 1.83*/

	//!Tipos de uso del println
	/*fmt.Println(nombre)
	fmt.Println(edad)
	fmt.Println(altura)*/

	//fmt.Println("El nombre es:", nombre, "La edad es:", edad, "La altura es:", altura)

	//!Declaracion multiple de variable (acepta autoasigancion)
	/*var pais, departamento, barrio = "Colombia", "Bogota", "Verbenal"
	fmt.Println(pais, departamento, barrio)*/

	//!Imprimir constante
	//fmt.Println(Curso)

	//!Operadores relacionales
	//* <, >, <=, >=, ==, != , Siempre se obtiene un booleano
	/*var edad = 10
	var resultado = edad > 5
	fmt.Println(resultado)*/

	//!Operadores logicos
	//* &&, ||, !
	/*resultado := 20 == 20 && 30 == 30 && (90 < 100 || 100 == 90)
	fmt.Println(resultado)*/

	//!Strings
	/*var curso string = "Curso de go profesional" //1
	var curso2 = "Curso de go"                   //2
	curso3 := "Curso"
	fmt.Println(curso, curso2, curso3)
	//* Longitud del texto
	longitud := len(curso3)
	fmt.Println(longitud)
	//* Acceder a posision del texto
	primerCaracter := curso[0]
	ultimoCaracter := curso[len(curso3)-1]
	fmt.Printf("%c\n", primerCaracter)
	fmt.Printf("%c\n", ultimoCaracter)*/

	//!Lectura de variables
	var nombre string
	var edad int
	var altura float32

	fmt.Print("Ingresa tu nombre: ")
	fmt.Scanf("%s", &nombre)
	fmt.Print("Ingresa tu edad: ")
	fmt.Scanf("%d", &edad)
	fmt.Print("Ingresa tu estatura: ")
	fmt.Scanf("%f", &altura)

	fmt.Printf("Hola %s con una edad %d y una altura %.2f\n", nombre, edad, altura)

}

// go build main.go --> ejecutable
// go run main.go --> ejecutar y compilar internamente
