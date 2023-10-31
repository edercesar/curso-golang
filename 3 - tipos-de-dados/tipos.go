package main

import "fmt"
import "errors"

func main()  {
	var  numero int16 = 90
	fmt.Println(numero)

	var erro error = errors.New("Meu erro")
	fmt.Println(erro)
}