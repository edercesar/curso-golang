package main

import "fmt"

func main() {
	fmt.Println("Aula de Ponteiros")
	var variavel1 int = 10
	var variavel2 int = variavel1

	fmt.Println(variavel1, variavel2)

	variavel1++
	fmt.Println(variavel1, variavel2)

	// ponteiro e uma referencia de memoria

	var variavel3 int = 12
	var ponteiro *int = &variavel3

	fmt.Println(variavel3, *ponteiro) // s agente remover o * ele vai printar o endereco de memoria da variavel 3
}
