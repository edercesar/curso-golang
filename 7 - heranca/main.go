package main

import "fmt"

type pessoa struct {
	nome string
	sobrenome string
	idade uint8
	altura uint8

}

type estudante  struct {
	pessoa
	curso string
	faculdade string
}

func main() {
	fmt.Println("Aula de heranca")

	pessoa1 := pessoa{"Joao", "Eduardo", 14, 150}
	fmt.Println(pessoa1)

	estudante1 := estudante{pessoa1, "Engenharia", "mackenzie"}
	fmt.Println(estudante1)
	fmt.Println("Idade do estudade",estudante1.idade)
	fmt.Println("Altura do estudante", estudante1.altura)
}