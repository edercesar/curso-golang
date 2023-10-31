package main

import "fmt"

type usuario struct {
	nome string
	idade uint8
	endereco end
}

type end struct {
	rua string
	numero uint32
}

func main() {
	fmt.Println("Arquivo structs")

	var u usuario
	u.nome = "Eder"
	u.idade = 39
	fmt.Println(u)
	enderecoexemplo := end{"Rua Cacique", 799}
	usuario2 := usuario{"mariana", 35, enderecoexemplo}
	fmt.Println(usuario2)
}