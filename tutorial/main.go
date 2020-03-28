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

	For1()
	For2()

	fmt.Println("<-------------->")
	fmt.Println("Bye bye, tutorial")
}
