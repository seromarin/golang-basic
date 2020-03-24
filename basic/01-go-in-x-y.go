package main

import "fmt"

func main() {
	var message string = "Hola mundo"
	easyMessage := "Hola mundo usando :="
	fmt.Println(message)
	fmt.Println(easyMessage)

	// Comments
	a := 10.
	var b float64 = 3
	fmt.Println(a / b)

	var c int = 10
	d := 3
	fmt.Println(c / d)

	x := true
	y := false

	fmt.Println(x || y)
	fmt.Println(x && y)
	fmt.Println(!x)

	privada()
	Publica()
	imprimirHola()
}

func privada() {
	fmt.Println("Ejecuta logica que no necesita ser exportada")
}

// Publica - This function can be used on other packages
func Publica() {
	fmt.Println("Ejecuta logica que quiero exportar")
}

func imprimirHola() {
	defer fmt.Println("Mundo")
	fmt.Println("Hola")
}
