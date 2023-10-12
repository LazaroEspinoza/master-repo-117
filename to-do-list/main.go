package main

import "fmt"

func main() {
	var a = [...]int{10, 20, 30, 40, 50}
	a[0] = 100
	a[1] = 200

	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}

	for index, value := range a {
		fmt.Printf("Indice: %d, Valor: %d\n", index, value)
	}
	var matriz = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(matriz)

	dias()
}

func dias() {
	diasSemana := []string{"Domingo", "Lunes", "Martes",
		"Miercoles", "Jueves", "Viernes", "Sabado"}

	diaRebanada := diasSemana[0:5]
	fmt.Println(diaRebanada)

	diaRebanada = append(diaRebanada, "Viernes", "Sabado")

	diaRebanada = append(diaRebanada[:2], diaRebanada[3:]...)
	fmt.Println(diaRebanada)

	fmt.Println(len(diaRebanada))
	fmt.Println(cap(diaRebanada))

	nombres := make([]string, 5, 10)
	nombres[0] = "lazaro"
	fmt.Println(nombres)

	rebanada1 := []int{1, 2, 3, 4, 5}
	rebanada2 := make([]int, 5)

	copy(rebanada1, rebanada2)

	fmt.Println(rebanada2)
	fmt.Println(rebanada1)

}
