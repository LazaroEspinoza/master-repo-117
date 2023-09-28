package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println("hola mundo")
	fmt.Println(quote.Glass())

	firstName := "Lazaro"
	lastName := "Espinoza"
	age := 26

	fmt.Println("my First name is", firstName, "\nMy Last Name is ", lastName, "\nI am ", age, "old")
}
