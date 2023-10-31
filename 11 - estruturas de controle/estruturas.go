package main

import "fmt"

func main() {
	numero := 99
	if numero > 50 {
		fmt.Println("o Numero  e maior que 50")
	} else {
		fmt.Println("o numero  e menor que 50")
	}
	if outronumero := numero; outronumero < 100 {
		fmt.Println("E  outro numero e menor que 100 ")
	} else {
		fmt.Println("E  outro numero e maior que 100 ")
	}
}
