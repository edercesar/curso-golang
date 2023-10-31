package main

import "fmt"

func fibonnaci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}
	return fibonnaci(posicao-2) + fibonnaci(posicao-1)

}

func main() {
	fmt.Println("Funcoes recursivas")
	posicao := uint(10)
	fmt.Println(fibonnaci(posicao))
}
