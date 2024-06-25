package animal

import "fmt"

type Animal interface {
	Sonido()
}

// estructura de perro y sus metodos
type Perro struct {
	Nombre string
}

func (p *Perro) Sonido() {
	fmt.Println(p.Nombre + " Hace guau guau")
}

//Estructura de gato y sus metodos
type Gato struct {
	Nombre string
}

func (g *Gato) Sonido() {
	fmt.Println(g.Nombre + " hace miau")
}

//Funcion para imprimir sonido
func HacerSonido(animal Animal) {
	animal.Sonido()
}
