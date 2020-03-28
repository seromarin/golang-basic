package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Hello, tutorial")
	fmt.Println("<-------------->")

	fmt.Printf("2 + 3 = %v\n", Add(2, 3))

	a, b := Swap("Hello", "World")
	fmt.Printf("Swap 'Hello' 'World' -> '%s' '%s'\n", a, b)

	c, d := Split(17)
	fmt.Printf("Split 17 = %v %v\n", c, d)

	fmt.Printf("El nuevo tipo de 15 es %v\n", reflect.TypeOf(Conversion(15)))

	fmt.Printf("El valor de PI es %v\n", Pi)

	// Cycles in Golang
	For1()
	For2()

	While()

	// Conditionals
	fmt.Println(Sqrt(2), Sqrt(-4))
	fmt.Println(
		Pow(3, 2, 10),
		Pow(3, 3, 20),
	)
	fmt.Println(
		Pow2(3, 2, 10),
		Pow2(3, 3, 20),
	)

	fmt.Println("<-------------->")
	fmt.Println("Bye bye, tutorial")
}
