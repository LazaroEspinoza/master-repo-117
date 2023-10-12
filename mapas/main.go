package main

import "fmt"

func main() {
	colors := map[string]string{
		"rojo":  "#FF0000",
		"verde": "#00FF00",
		"azul":  "#0000FF",
	}
	fmt.Println(colors["rojo"])
	colors["negro"] = "#000000"

	fmt.Println(colors)
	valor, ok := colors["rojo"]
	if ok {
		fmt.Println("Si existe la clave")
	} else {
		fmt.Println("No existe la clave")
	}

	fmt.Println(valor, ok)

	delete(colors, "azul")

	fmt.Println(colors)
	colors["azul"] = "#0000FF"

	for clave, valor := range colors {
		fmt.Printf("clave: %s , Valor: %s\n", clave, valor)
	}

}
