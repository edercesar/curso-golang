package  main

import "fmt"

func main()  {
	// ARITIMETICOS
	soma := 1 + 2
	subtracao := 2-1
	divisao := 10/2
	multiplicacao := 4*2
	restodadivisao := 10 % 3

	fmt.Println(soma, subtracao, divisao, multiplicacao, restodadivisao)

	//ATRIBUICOES

	var variavel string = "String1"
	variavel2 := "String 2"

	fmt.Println(variavel, variavel2)

	//OPERADORES RELACIONAIS

	fmt.Println(1 > 2, 2 > 1, 3 ==3)

	//OPERADORES LOGICOS

	fmt.Println("--------------")
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro) // negando o valor
	fmt.Println(!falso) // negando o valor

	//OPERADORES UNARIOS

	numero := 15
	numero++ //incrementa de 1 em 1
	numero += 10 //incrementa de 10 em 10 e o valor que vc quiser 
	fmt.Println(numero)

	numero-- // decrementa de 1 em 1
	numero-=8 //decrementa de 8 em 8 e o numero que vc quiser

	numero *=2 //multiplica o valor da variavel por 2
	numero /=2  // faz a divisao do valor da variavel por 2

	fmt.Println(numero)

	//condicional ternario
	var texto string
	if numero > 10 {
		texto = "numero e maior que 10"
	} else {
		texto = "numeo e menor ou igual a 10"
	}
	

	fmt.Println(texto)
}

