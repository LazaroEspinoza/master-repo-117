package main

import (
	"fmt"

	"github.com/MASTER-REPO-117/library/animal"
	"github.com/MASTER-REPO-117/library/book"
)

func main() {
	myBook := book.NewBook("Moby Dick", "Herman Melvine", 300)

	var milibro = book.Book{
		Title:  "Alicia en el pais de las maravillas",
		Author: "Lewis Carroll",
		Pages:  208,
	}

	milibro.PrintInfo()
	myTextBook := book.NewTextBook("geografia", "juan chavez", 322, "santillansa eos", "secundaria")

	myTextBook.PrintInfo()
	myBook.SetTittle("Moby Dick (Special Edition)")
	fmt.Println(myBook.GetTitle())
	myTexBook := book.NewTexBook("Comunicacion", "Luis Aguerri", 684, "Composiciones literarias ROSS",
		"Primer Grado de Maestria")

	//myBook.PrintInfo()
	//myTexBook.PrintInfo()

	book.Print(myBook)
	book.Print(myTexBook)

	animales := []animal.Animal{
		&animal.Perro{Nombre: "Pelusa"},
		&animal.Gato{Nombre: "malteada"},
		&animal.Perro{Nombre: "Zeus"},
		&animal.Gato{Nombre: "Edu"},
	}

	for _, animal := range animales {
		animal.Sonido()
	}

}
