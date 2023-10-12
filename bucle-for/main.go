package main

import "fmt"

func main() {

	for i := 1; i <= 10; i++ {
		fmt.Println(i)
		if i == 5 {
			break

		}
	}

	for e := 1; e <= 10; e++ {
		if e == 5 {
			continue
		}
		fmt.Println(e)

	}
	saludo := holaa("Patricio")
	fmt.Println(saludo)

	suma, multi := calc(4, 9)
	fmt.Println("La suma es: ", suma)
	fmt.Println("La multiplicacion es: ", multi)

}

func holaa(name string) string {
	return "Hola, " + name
}

func calc(a, b int) (suma, multi int) {
	suma = a + b
	multi = a * b
	return
}
