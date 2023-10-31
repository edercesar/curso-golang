package main

import (
	"fmt"
	"modulo/auxiliar"
	"github.com/badoux/checkmail"
) 


func main()  {
	fmt.Println("Escrevendo do arquivo main")
	auxiliar.Escrever()
	auxiliar.Escrever2()
	erro := checkmail.ValidateFormat("eder.luz@ifood.com.br")
	fmt.Println(erro)
}
