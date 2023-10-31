package main

import "fmt"

//funcao com dois retornos
func calculos(numero1, numero2 int8) (int8, int8) {
	soma := numero1 + numero2
	subtracao := numero1 - numero2
	return soma, subtracao

}

//funcao com um retorno
func somar(n1, n2 int8) int8 {
	return n1 + n2

}

func main() {
	soma := somar(10, 30)
	fmt.Println(soma)

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	resultado := f("Texto da funçao 1")
	fmt.Println(resultado)

	//printando os dois retornos
	result_soma, resul_sub := calculos(10, 15)
	fmt.Println(result_soma, resul_sub)

	//printando somente a soma
	printa_soma, _ := calculos(10, 15)
	fmt.Println(printa_soma)

	//printando somente a subtracao
	_, printa_sub := calculos(10, 15)
	fmt.Println(printa_sub)
}
