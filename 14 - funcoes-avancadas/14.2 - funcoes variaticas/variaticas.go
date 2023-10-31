package main

import "fmt"

func calculos(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}

	return total
}

func TesteStringcomInt(texto string, numeros ...int) {
	fmt.Println(texto, numeros)
}

func main() {
	soma := calculos(1, 4, 6, 7, 7, 2, 100, 3000, 43)
	fmt.Println(soma)

	TesteStringcomInt("Ola terraquios", 2, 6, 6, 8, 10)
}
