package main

import "fmt"

func main() {
	func() {
		fmt.Println("Ola mudo")
	}()

	teste := func(texto string) string {
		return fmt.Sprintf("Parametro passado > %s", texto)
	}("Eu sou o Parametro")

	fmt.Println(teste)
}
