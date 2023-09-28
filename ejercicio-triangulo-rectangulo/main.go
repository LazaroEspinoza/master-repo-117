package main

import (
	"fmt"
	"math"
)

func main() {
	var lado1, lado2 float64
	const precision = 2

	fmt.Println("Ingrese la medida de el lado 1 del triangulo:")
	fmt.Scanln(&lado1)
	fmt.Println("Ingrese la medida de el lado 2 del triangulo:")
	fmt.Scanln(&lado2)

	area := (lado1 * lado2) / 2
	hipotenusa := math.Sqrt(math.Pow(lado1, 2) + math.Pow(lado2, 2))
	perimetro := lado1 + lado2 + hipotenusa

	fmt.Printf("\nEl area es: %.*f \n\n", precision, area)
	fmt.Printf("El perimetro es: %.*f \n", precision, perimetro)

	x := 5
	y := 10
	z := 15

	resultado := ((x+y)*z)/(x*y) > z && x != y

	fmt.Println(resultado)
}
