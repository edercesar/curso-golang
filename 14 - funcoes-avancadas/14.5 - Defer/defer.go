package main

import "fmt"

func funcao1() {
	fmt.Println("Executando a funcao 1")
}

func funcao2() {
	fmt.Println("Executando a funcao 2")
}

func MediaEscolar(n1, n2 float32) bool {
	media := (n1 + n2) / 2
	if media >= 6 {
		return true
	}
	return false
}
func main() {
	fmt.Println(MediaEscolar(9, 5))
	defer funcao1()
	funcao2()
	fmt.Println("Eu sou o ultimo Print da funcao main")
}
