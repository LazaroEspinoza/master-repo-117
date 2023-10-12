package main

import "fmt"

type Persona struct {
	nombre string
	edad   int
	correo string
}

func main() {
	var p Persona
	p.nombre = "Lazaro"
	p.edad = 26
	p.correo = "master_chief1237@hotmail.com"

	fmt.Println(p)

	l := Persona{"Patricio", 26, "lazarochief1237@gmail.com"}

	fmt.Println(l)

	l2 := Persona{"Juan", 45, "juanguerra@gmail.com"}

	fmt.Println(l2)

	var x int = 10
	fmt.Println(x)
	editar(&x)
	fmt.Println(x)

}
func editar(x *int) {
	*x = 20
}
