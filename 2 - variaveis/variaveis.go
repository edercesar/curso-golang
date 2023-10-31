package main

import (
	"fmt"
)

func main()  {
	var variavel string = "Sou uma String"
	variavel2 := "Eu tbm sou uma string implicita"
	variavel3 := 102030
	variavel4:= true

	fmt.Println(variavel, variavel2, variavel3, variavel4)

	var (
		variavel5 string = "Eu sou a variavel 5"
		variavel6 string = "Eu sou a variavel 6"	
	)

	variavel7, variavel8 := "Eu sou a variavel 7"," e eu sou a variavel 8"
	fmt.Println(variavel5, variavel6, variavel7, variavel8 )

	const constante1 string = "Eu sou uma  constante"
	fmt.Println(constante1)
}