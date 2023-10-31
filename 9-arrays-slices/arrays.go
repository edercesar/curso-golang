package main

import "fmt"

func main() {
	fmt.Println("Aula de Arrays e Slices")

	var array1 [5]string    //array declarado com 5 posicoes
	array1[0] = "Posicao 1" //definido o valor 'posicao 1' para o indice 0 do array
	fmt.Println(array1)

	//declaracao de array com 5 posicoes e ja declarando os valores
	array2 := [5]string{"Posicao 1", "Posicao 2", "Posicao 3", "Posicao 4", "Posicao 5"}
	fmt.Println(array2)

	//neste caso a quantidade de arrays e definida pela quantidade de valores que voce declarar
	// e nao tem possbilidade de incrementar depois
	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	// Slices
	slice := []int{10, 11, 12, 13, 14, 15}
	fmt.Println(slice)

	slice = append(slice, 16)
	fmt.Println(slice)

	slice2 := array3[1:3]
	fmt.Println(slice2)

	array2[1] = "Posicao Alterada"
	fmt.Println(array2)

	// Arrays internos
	fmt.Println("------------- ARRAYS INTERNOS ---------")
	slice3 := make([]float32, 10, 11)
	fmt.Print(len(slice3))   // para ver o tamanho do array que e o 10
	fmt.Println(cap(slice3)) // para ver a capacidade do array que e 11
	fmt.Println(slice3)

	//slice sem declarar a capacidade
	slice4 := make([]float32, 5)
	fmt.Print(len(slice4))     // para ver o tamanho do array que e o 5
	fmt.Println(cap(slice4))   // para ver a capacidade do array que vai ser igual o tamanho no caso 5
	slice4 = append(slice4, 6) // apendando um novo valor no slice
	fmt.Print(len(slice4))     // como apendei um novo valor agora o tamanho dele vai ser6
	fmt.Println(cap(slice4))   // agora como ele tem o tamanho de 6, a capacidade vai dobrar para 12

}
