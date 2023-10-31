package main

import "fmt"

func main() {
	// i := 0
	// for i < 10 {
	// 	i++
	// 	fmt.Println("Incrementando i")
	// 	time.Sleep(time.Second)
	// }
	// fmt.Println("Jeito de fazer o loop 1", i)

	// for j := 0; j < 10; j++ {
	// 	fmt.Println("Incrementando J", j)
	// 	time.Sleep(time.Second)
	// }

	nomes := [3]string{"Joao", "Rafael", "Mariana"}
	for _, nome := range nomes {
		fmt.Println(nome)
	}

	for i, palavra := range "PALAVRA" {
		fmt.Println(i, string(palavra)) // se nao colocar a funcao string em "palavra ele vai trazer o valor de cada letra da tabel ascii"
	}

	usuario := map[string]string{
		"nome":      "Eder",
		"sobrenome": "Luz",
	}
	for _, valor := range usuario {
		fmt.Println(valor)
	}

}
