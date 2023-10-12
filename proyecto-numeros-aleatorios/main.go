package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println(rand.Intn(100))
	juego()
}

func juego() {
	numAleatorio := rand.Intn(100)
	var numIngresado int
	var intentos int
	const maxIntentos = 10

	for intentos < maxIntentos {
		intentos++
		fmt.Printf("Ingrese un numero (intentos restantes: %d) ", maxIntentos-intentos+1)
		fmt.Scanln(&numIngresado)

		if numIngresado == numAleatorio {
			fmt.Println("Felicitaciones, adivinaste el numero!!!")
			jugarNuevamente()
			return
		} else if numIngresado < numAleatorio {
			fmt.Println("El numero a Adivinar es mayor.")
		} else if numIngresado > numAleatorio {
			fmt.Println("El numero a Adivinar es menor.")
		}
	}
	fmt.Println("Se acabaron los intentos :(( . El numero a adivinar era: ", numAleatorio)
	jugarNuevamente()
}
func jugarNuevamente() {
	var eleccion string
	fmt.Print("Quieres jugar nuevamente? (s/n): ")
	fmt.Scanln(&eleccion)

	switch eleccion {
	case "s":
		juego()
	case "n":
		fmt.Println("--> Gracias por jugar!!!")
	default:
		fmt.Println("Eleccion invalida. Intentalo nuevamente.")
		jugarNuevamente()
	}

}
